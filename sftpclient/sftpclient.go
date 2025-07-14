package sftpclient

import (
	"io"
	"os"

	"github.com/pkg/sftp"
)

type SFTPClient interface {
	ReadDir(path string) ([]os.FileInfo, error)
	Open(path string) (File, error)
}

type File interface {
	io.Reader
	io.Closer
}

type RealSFTPClient struct {
	Client *sftp.Client
}

func (r *RealSFTPClient) ReadDir(path string) ([]os.FileInfo, error) {
	return r.Client.ReadDir(path)
}

func (r *RealSFTPClient) Open(path string) (File, error) {
	return r.Client.Open(path)
}
