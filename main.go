package main

import (
	"log"
	"os"
	"sync"
	"time"

	"fileparser/domain"
	"fileparser/downloader"
	"fileparser/extract"
	"fileparser/hasher"
	"fileparser/indexer"
	"fileparser/sftpclient"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func connectSFTP() (*sftp.Client, error) {
	config := &ssh.ClientConfig{
		User:            "devuser",
		Auth:            []ssh.AuthMethod{ssh.Password("devpass")},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // Dev only
	}
	conn, err := ssh.Dial("tcp", "sftp-server:22", config)
	if err != nil {
		return nil, err
	}
	return sftp.NewClient(conn)
}

func main() {
	os.MkdirAll("downloads", 0755)

	sftpClient, err := connectSFTP()
	if err != nil {
		log.Fatalf("SFTP connection error: %v", err)
	}
	defer sftpClient.Close()

	es_client, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Elasticsearch error: %v", err)
	}

	es := &indexer.ElasticAdapter{Client: es_client}

	filePathsChan := make(chan string, 10)
	fileDataChan := make(chan domain.FileData, 10)
	hashedDataChan := make(chan domain.FileData, 10)

	var downloadWG, readWG, indexWG, hashWG sync.WaitGroup

	downloadWG.Add(1)
	go downloader.DownloadFiles(
		&sftpclient.RealSFTPClient{Client: sftpClient},
		filePathsChan,
		&downloadWG,
	)

	extractWorkers := 4
	transformWorkers := 4
	loadWorkers := 4

	//EXTRACT
	readWG.Add(extractWorkers)
	for i := 0; i < extractWorkers; i++ {
		go extract.ReadFiles(filePathsChan, fileDataChan, &readWG)
	}

	//TRANSFORM
	hashWG.Add(transformWorkers)
	for i := 0; i < transformWorkers; i++ {
		go hasher.HashFiles(fileDataChan, hashedDataChan, &hashWG)
	}

	//LOAD
	indexWG.Add(loadWorkers)
	for i := 0; i < loadWorkers; i++ {
		go indexer.IndexFiles(es, "sftp_index", hashedDataChan, &indexWG)
	}

	go func() {
		downloadWG.Wait()
		close(filePathsChan)
	}()

	go func() {
		readWG.Wait()
		close(fileDataChan)
	}()

	go func() {
		hashWG.Wait()
		close(hashedDataChan)
	}()

	indexWG.Wait()

	log.Println("✅ All files indexed. Waiting 10 minutes before terminating")
	time.Sleep(10 * time.Minute)
}
