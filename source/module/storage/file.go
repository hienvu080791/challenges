package storage

import (
	"fmt"
	"sync"
	"time"
)

var fileStore = sync.Map{}

type FileMetadata struct {
	Filename    string    `json:"filename"`
	ContentType string    `json:"content_type"`
	Size        int64     `json:"size"`
	TempPath    string    `json:"temp_path"`
	UploadTime  time.Time `json:"upload_time"`
	RequestInfo string    `json:"request_info"`
}

func CreateFileMetadata(metadata FileMetadata) error {
	_, isExist := fileStore.LoadOrStore(metadata.Filename, metadata)
	if isExist {
		return fmt.Errorf("file %s already exists. Please try another one", metadata.Filename)
	}
	//fmt.Println("Current file in fileStore:")
	//fileStore.Range(func(key, value interface{}) bool {
	//	fileMeta := value.(FileMetadata)
	//	fmt.Printf("Filename: %s, ContentType: %s, Size: %d, TempPath: %s, UploadTime: %s, RequestInfo: %s\n", fileMeta.Filename, fileMeta.ContentType, fileMeta.Size, fileMeta.TempPath, fileMeta.UploadTime.Format(time.RFC3339), fileMeta.RequestInfo)
	//	return true
	//})
	return nil
}
