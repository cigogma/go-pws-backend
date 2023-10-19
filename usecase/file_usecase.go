package usecase

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"time"

	"cloud.google.com/go/storage"
	"github.com/gabriel-vasile/mimetype"
	"github.com/google/uuid"
	"github.com/pws-backend/domain"
)

type fileUsecase struct {
	fileRepository domain.FileRepository
	contextTimeout time.Duration
}

func (fu *fileUsecase) Delete(c context.Context, file *domain.File) error {
	client, err := storage.NewClient(c)
	if err != nil {
		return fmt.Errorf("storage.NewClient: %w", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(c, time.Second*30)
	defer cancel()

	o := client.Bucket(file.Bucket).Object(file.Key)
	o.Delete(ctx)

	return fu.fileRepository.Delete(c, file)
}

func (*fileUsecase) Upload(c context.Context, file multipart.File, fileE *domain.File) error {

	bucket := "pws-frontend"
	fileExtension, err := mimetype.DetectReader(file)

	if err != nil {
		return err
	}

	object := "go-test/" + uuid.New().String() + fileExtension.Extension()

	client, err := storage.NewClient(c)
	if err != nil {
		return fmt.Errorf("storage.NewClient: %w", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(c, time.Second*30)
	defer cancel()

	o := client.Bucket(bucket).Object(object)

	o = o.If(storage.Conditions{DoesNotExist: true})
	file.Seek(0, io.SeekStart)
	wc := o.NewWriter(ctx)
	if _, err = io.Copy(wc, file); err != nil {
		return fmt.Errorf("io.Copy: %w", err)
	}
	if err := wc.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %w", err)
	}
	fileE.Bucket = bucket
	fileE.Key = object
	fileE.Mime = fileExtension.String()
	fileE.StorageClass = "google"
	return nil
}

func (fu *fileUsecase) Create(c context.Context, file *domain.File) error {
	return fu.fileRepository.Create(c, file)
}

func (fu *fileUsecase) Fetch(c context.Context, limit int) ([]domain.File, error) {
	return fu.fileRepository.Fetch(c, limit)
}

func (fu *fileUsecase) GetByID(c context.Context, fileId uint) (domain.File, error) {
	return fu.fileRepository.GetByID(c, fileId)
}

func NewFileUsecase(fileRepository domain.FileRepository, timeout time.Duration) domain.FileUsecase {
	return &fileUsecase{
		fileRepository: fileRepository,
		contextTimeout: timeout,
	}
}
