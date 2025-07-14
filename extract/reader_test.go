package extract

import (
	"fileparser/domain"
	"os"
	"path/filepath"
	"sync"
	"testing"
)

func TestReadFiles(t *testing.T) {
	// Create temporary file
	tmpFile, err := os.CreateTemp("", "testfile.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	//Write names to the temp file
	lines := []string{"Syed Hart", "Kane Dunlap", "Isobella Dickerson"}
	for _, line := range lines {
		_, err := tmpFile.WriteString(line + "\n")
		if err != nil {
			t.Fatalf("Failed to write to temp file: %v", err)
		}
	}
	tmpFile.Close()

	// Prepare channel to push file names to and publish data out
	in := make(chan string, 1)
	out := make(chan domain.FileData, 10)
	var wg sync.WaitGroup

	in <- tmpFile.Name()
	close(in)

	wg.Add(1)
	go ReadFiles(in, out, &wg)

	wg.Wait()
	close(out)

	var results []domain.FileData
	for fd := range out {
		results = append(results, fd)
	}

	if len(results) != len(lines) {
		t.Errorf("Expected %d lines, got %d", len(lines), len(results))
	}

	for i, fd := range results {
		if fd.Content != lines[i] {
			t.Errorf("Expected line %q, got %q", lines[i], fd.Content)
		}
		if filepath.Base(fd.Name) != filepath.Base(tmpFile.Name()) {
			t.Errorf("Unexpected name in output: %s", fd.Name)
		}
	}
}
