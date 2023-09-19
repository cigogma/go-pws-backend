package domain

import "gorm.io/gorm"

type File struct {
	gorm.Model

	Key          string
	Bucket       string
	Mime         string
	StorageClass string
	Comment      string
}

// Define a function to retrieve the URL for a File
func (file *File) GetURL() string {
	return "https://storage.googleapis.com/" + file.Bucket + "/" + file.Key
}
