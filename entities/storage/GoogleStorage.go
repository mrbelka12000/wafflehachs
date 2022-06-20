package storage

import (
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"time"

	"cloud.google.com/go/storage"
)

const GoogleConfigFileName = "entities/storage/gcp-storage.json"

type ClientUploader struct {
	cl         *storage.Client
	projectID  string
	bucketName string
	uploadPath string
}

func getClient() *ClientUploader {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", GoogleConfigFileName) // FILL IN WITH YOUR FILE PATH
	client, err := storage.NewClient(context.Background())
	if err != nil {
		return nil
	}

	return &ClientUploader{
		cl:         client,
		bucketName: os.Getenv("bucket"),
		projectID:  os.Getenv("gcp-storage-353811"),
		uploadPath: os.Getenv("uploadPath") + "/",
	}
}

// UploadFile uploads an object
func UploadFile(file multipart.File, object string) error {
	c := getClient()
	if c == nil {
		return errors.New("не удалось получить клиента для загрузки файлов")
	}

	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	// Upload an object with storage.Writer.

	wc := c.cl.Bucket(c.bucketName).Object(c.uploadPath + object).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %v", err)
	}

	return nil
}

func DeleteFile(object string) error {
	// bucket := "bucket-name"
	// object := "object-name"
	ctx := context.Background()
	c := getClient()
	if c == nil {
		return errors.New("Клиента нет")
	}
	defer c.cl.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	o := c.cl.Bucket(c.bucketName).Object(object)

	// Optional: set a generation-match precondition to avoid potential race
	// conditions and data corruptions. The request to upload is aborted if the
	// object's generation number does not match your precondition.
	attrs, err := o.Attrs(ctx)
	if err != nil {
		return fmt.Errorf("object.Attrs: %v", err)
	}
	o = o.If(storage.Conditions{GenerationMatch: attrs.Generation})

	if err := o.Delete(ctx); err != nil {
		return fmt.Errorf("Object(%q).Delete: %v", object, err)
	}
	return nil
}
