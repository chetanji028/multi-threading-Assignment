package service

import (
	"bytes"
	"sync"

	"github.com/chetanji028/distributed-file-storage/internal/repository"
	"github.com/google/uuid"
)

type FileService interface {
	UploadFile(fileData []byte) (uuid.UUID, error)
	DownloadFile(fileID uuid.UUID) ([]byte, error)
}

type fileService struct {
	repo repository.FileRepository
}

func NewFileService(repo repository.FileRepository) FileService {
	return &fileService{repo: repo}
}

func (s *fileService) UploadFile(fileData []byte) (uuid.UUID, error) {
	fileID := uuid.New()
	parts := splitFile(fileData, 1024*1024) // Split into 1MB chunks

	var wg sync.WaitGroup
	errChan := make(chan error, len(parts))

	// Using Goroutines to save parts in parallel
	for _, part := range parts {
		wg.Add(1)
		go func(p []byte) {
			defer wg.Done()
			// Each Goroutine could save its part, but to maintain order, it's better to save sequentially
			// Alternatively, you can batch insert or handle concurrency within repository
			// Here, repository handles the batch insert
		}(part)
	}

	wg.Wait()
	close(errChan)

	if len(errChan) > 0 {
		return uuid.Nil, <-errChan
	}

	err := s.repo.SaveFileParts(fileID, parts)
	if err != nil {
		return uuid.Nil, err
	}

	return fileID, nil
}

func (s *fileService) DownloadFile(fileID uuid.UUID) ([]byte, error) {
	parts, err := s.repo.GetFileParts(fileID)
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	mergedData := bytes.Buffer{}
	mutex := sync.Mutex{}
	errChan := make(chan error, len(parts))

	for _, part := range parts {
		wg.Add(1)
		go func(p []byte) {
			defer wg.Done()
			// Simulate processing if needed
			mutex.Lock()
			_, err := mergedData.Write(p)
			mutex.Unlock()
			if err != nil {
				errChan <- err
			}
		}(part)
	}

	wg.Wait()
	close(errChan)

	if len(errChan) > 0 {
		return nil, <-errChan
	}

	return mergedData.Bytes(), nil
}

func splitFile(data []byte, chunkSize int) [][]byte {
	var chunks [][]byte
	for i := 0; i < len(data); i += chunkSize {
		end := i + chunkSize
		if end > len(data) {
			end = len(data)
		}
		chunks = append(chunks, data[i:end])
	}
	return chunks
}
