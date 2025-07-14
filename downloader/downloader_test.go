package downloader_test

import (
	"fmt"
	"io"
	"os"
	"sync"
	"testing"
	"time"

	"fileparser/downloader"
	"fileparser/mocks"

	"go.uber.org/mock/gomock"
)

type fakeFileInfo struct {
	name string
}

func (f fakeFileInfo) IsDir() bool        { return false }
func (f fakeFileInfo) Sys() interface{}   { return nil }
func (f fakeFileInfo) Name() string       { return f.name }
func (f fakeFileInfo) Size() int64        { return 123 }
func (f fakeFileInfo) Mode() os.FileMode  { return 0644 }
func (f fakeFileInfo) ModTime() time.Time { return time.Now() }

func TestDownloadFiles(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mocks.NewMockSFTPClient(ctrl)
	mockFile := mocks.NewMockFile(ctrl)

	dir := "upload/"
	filename := "test.txt"
	content := "mock content"

	mockClient.EXPECT().
		ReadDir(dir).
		Return([]os.FileInfo{fakeFileInfo{name: filename}}, nil)

	mockClient.EXPECT().
		Open(fmt.Sprintf("%s%s", dir, filename)).
		Return(mockFile, nil)

	mockFile.EXPECT().
		Read(gomock.Any()).
		DoAndReturn(func(p []byte) (int, error) {
			copy(p, content)
			return len(content), io.EOF
		})

	mockFile.EXPECT().Close().Return(nil)

	ch := make(chan string, 1)
	var wg sync.WaitGroup
	wg.Add(1)

	go downloader.DownloadFiles(mockClient, ch, &wg)
	wg.Wait()
	close(ch)

	path := <-ch
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("Failed to read downloaded file: %v", err)
	}

	if string(data) != content {
		t.Errorf("Expected content %q, got %q", content, string(data))
	}

	os.RemoveAll("downloads")
}
