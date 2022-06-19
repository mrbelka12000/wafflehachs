package storage

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"time"

	"cloud.google.com/go/storage"
)

const GoogleConfigFileName = "entities/storage/gcp-storage.json"

// import (
// 	"context"
// 	"fmt"
// 	"io"
// 	"log"
// 	"mime/multipart"
// 	"net/http"
// 	"os"
// 	"time"
// 	"wafflehacks/tools"

// 	"cloud.google.com/go/storage"
// 	"github.com/gin-gonic/gin"
// )

type ClientUploader struct {
	cl         *storage.Client
	projectID  string
	bucketName string
	uploadPath string
}

// var uploader *ClientUploader

// func init() {
// 	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "check.json") // FILL IN WITH YOUR FILE PATH
// 	client, err := storage.NewClient(context.Background())
// 	if err != nil {
// 		log.Fatalf("Failed to create client: %v", err)
// 	}

// 	uploader = &ClientUploader{
// 		cl:         client,
// 		bucketName: bucketName,
// 		projectID:  projectID,
// 		uploadPath: "test-files/",
// 	}
// }

// func main() {
// 	// uploader.UploadFile("notes_test/abc.txt")
// 	r := gin.Default()
// 	r.GET("/", func(c *gin.Context) {
// 		c.Header("Content/type", "text/html")
// 		c.Writer.Write([]byte(`<!DOCTYPE html>
// 		<html lang="en">
// 		<head>
// 			<meta charset="UTF-8">
// 			<meta name="viewport" content="width=device-width, initial-scale=1.0">
// 			<meta http-equiv="X-UA-Compatible" content="ie=edge">
// 			<link rel="stylesheet" type="text/css" href="assets/css/createpost.css">
// 			<link rel="shortcut icon" href="#">
// 			<title>Create Post</title>
// 		</head>
// 		<body>
// 			<div class="title">Create post</div>
// 			<form method="POST"enctype=multipart/form-data action="/upload" >
// 				<div class="container">

// 					<input type="file" accept=".jpg, .jpeg, .png , .gif"name="file_input"/>

// 					</div>
// 					<br>

// 					<div class="btn_container" style="background-color:#f1f1f1">
// 						<button class="btn" type="submit">Create</button>
// 						<input class="btn" type="button" value="Home" onClick='location.href="http://localhost:8080/"'>
// 					</div>
// 				</div>
// 			</form>
// 		</body>
// 		</html>
// 		`))
// 	})
// 	r.POST("/upload", func(c *gin.Context) {
// 		f, err := c.FormFile("file_input")
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{
// 				"error": err.Error(),
// 			})
// 			return
// 		}

// 		blobFile, err := f.Open()
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{
// 				"error": err.Error(),
// 			})
// 			return
// 		}

// 		err = uploader.UploadFile(blobFile, tools.GetRandomString())
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{
// 				"error": err.Error(),
// 			})
// 			return
// 		}

// 		c.JSON(200, gin.H{
// 			"message": "success",
// 		})
// 	})

// 	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
// }

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
