package indexer_test

import (
	"fileparser/domain"
	"fileparser/indexer"
	"sync"
	"testing"

	"fileparser/mocks"

	"go.uber.org/mock/gomock"
)

func TestIndexFiles_WithGoMock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockIndexService(ctrl)
	mockClient := mocks.NewMockESClient(ctrl)

	file := domain.FileData{
		Name:    "test.txt",
		Content: "mocked content",
	}

	expectedDoc := map[string]string{
		"filename": file.Name,
		"content":  file.Content,
	}

	// Expectations
	mockClient.EXPECT().
		Index().
		Return(mockService)

	mockService.EXPECT().
		Index("test-index").
		Return(mockService)

	mockService.EXPECT().
		BodyJson(expectedDoc).
		Return(mockService)

	mockService.EXPECT().
		Do(gomock.Any()).
		Return(&indexer.IndexResponse{ID: "1"}, nil)

	// Run IndexFiles
	fileChan := make(chan domain.FileData, 1)
	fileChan <- file
	close(fileChan)

	var wg sync.WaitGroup
	wg.Add(1)

	indexer.IndexFiles(mockClient, "test-index", fileChan, &wg)
	wg.Wait()
}
