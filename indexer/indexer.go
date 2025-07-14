package indexer

import (
	"context"
	"fileparser/domain"
	"log"
	"sync"
)

func IndexFiles(esClient ESClient, indexName string, in <-chan domain.FileData, wg *sync.WaitGroup) {
	defer wg.Done()

	ctx := context.Background()

	for file := range in {
		doc := map[string]string{
			"filename": file.Name,
			"content":  file.Content,
		}

		_, err := esClient.Index().
			Index(indexName).
			BodyJson(doc).
			Do(ctx)

		if err != nil {
			log.Printf("Failed to index document %s: %v", file.Content, err)
		} else {
			log.Printf("Indexed from file %s row value %s into index %s", file.Name, file.Content, indexName)
		}
	}
}
