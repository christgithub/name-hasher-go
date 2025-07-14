package extract

import (
	"bufio"
	"fileparser/domain"
	"log"
	"os"
	"sync"
)

func ReadFiles(in <-chan string, out chan<- domain.FileData, wg *sync.WaitGroup) {
	defer wg.Done()
	for path := range in {
		log.Printf("Read file : %s", path)
		file, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			out <- domain.FileData{Name: path, Content: scanner.Text()}
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}
