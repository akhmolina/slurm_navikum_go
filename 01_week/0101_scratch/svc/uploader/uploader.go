package uploader

import (
	"context"
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

type Svc struct {
}

func NewSvc() *Svc {
	return &Svc{}
}

func (s *Svc) Upload(ctx context.Context, header *multipart.FileHeader, input multipart.File) (err error) {
	output, err := os.Create(filepath.Join("/upload", uuid.NewString()))
	if err != nil {
		return
	}

	n, err := io.Copy(output, input)
	if err != nil {
		return
	}

	if n != header.Size {
		err = errors.New("can't upload file, size mismatch")
		return
	}
	return
}
