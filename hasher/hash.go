package hasher

import (
	"crypto/sha256"
	"encoding/hex"
	"fileparser/domain"
	"log"
	"sync"
)

func HashFiles(in <-chan domain.FileData, out chan<- domain.FileData, wg *sync.WaitGroup) {
	defer wg.Done()
	for file := range in {
		hash := sha256.Sum256([]byte(file.Content))
		hashed := hex.EncodeToString(hash[:])
		out <- domain.FileData{Name: file.Name, Content: hashed}
		log.Printf("🔐 Hashed content from file %s", file.Name)
	}
}
