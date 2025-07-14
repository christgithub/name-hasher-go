package downloader

import (
	"fileparser/sftpclient"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
)

func DownloadFiles(client sftpclient.SFTPClient, out chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	uploadDir := "upload/"
	files, err := client.ReadDir(uploadDir)

	if err != nil {
		log.Printf("Failed to read from directory %s", err)
	}

	log.Printf("File in upload dir %v\n", len(files))

	for _, f := range files {
		name := f.Name()
		log.Printf("Preparing to download file %s\n", name)
		src, err := client.Open(filepath.Join(uploadDir, name))
		if err != nil {
			log.Printf("Failed to open file %s", name)
		}
		defer src.Close()

		os.MkdirAll("downloaded", 0755)
		dstPath := filepath.Join("downloaded/", name)
		dst, err := os.Create(dstPath)
		if err != nil {
			log.Printf("Failed to create local file %s", dst.Name())
		}

		io.Copy(dst, src)
		dst.Close()

		out <- dstPath
	}
}
