package handler

import (
	"fmt"
	"net/http"
	"wafflehacks/entities/storage"
	"wafflehacks/tools"
)

func (h *Handler) Upload(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(500)
		return
	}
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		SendErrorResponse(w, "Слишком большой файл", 400)
		h.log.Debug("Файл слишком много весит")
		return
	}
	file, handler, err := r.FormFile("avatar")
	if err != nil {
		SendErrorResponse(w, "Не удалось получить файл", 400)
		h.log.Debug("Не удалось получить файл по причине: " + err.Error())
		return
	}
	if handler.Size >= 20000000 {
		SendErrorResponse(w, "Изображение весит больше чем положено, 20mb", 400)
		h.log.Debug("Слишком большой файл")
		return
	}
	defer file.Close()

	if fileType, ok := tools.IsValidType(file); !ok {
		SendErrorResponse(w, fmt.Sprintf("Разрешение %v не поддерживается", fileType), 400)
		h.log.Debug(fmt.Sprintf("Разрешение %v не поддерживается", fileType))
		return
	}

	file.Seek(0, 0)

	filename := tools.GetRandomString()
	if err := storage.UploadFile(file, filename); err != nil {
		SendErrorResponse(w, "Не удалось загрузить файл ", 500)
		h.log.Debug("Не удалось загрузить файл по причине: " + err.Error())
		return
	}
	h.log.Info(tools.GetStorageUrl(filename))
	w.Write([]byte(tools.GetStorageUrl(filename)))
	// http.Redirect(w, r, tools.GetStorageUrl(filename), 301)
}

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

// const (
// 	projectID  = "gcp-storage-353811" // FILL IN WITH YOURS
// 	bucketName = "wafflehacksbucket"  // FILL IN WITH YOURS
// )

// type ClientUploader struct {
// 	cl         *storage.Client
// 	projectID  string
// 	bucketName string
// 	uploadPath string
// }

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

// // UploadFile uploads an object
// func (c *ClientUploader) UploadFile(file multipart.File, object string) error {
// 	ctx := context.Background()

// 	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
// 	defer cancel()

// 	// Upload an object with storage.Writer.

// 	wc := c.cl.Bucket(c.bucketName).Object(c.uploadPath + object).NewWriter(ctx)
// 	if _, err := io.Copy(wc, file); err != nil {
// 		return fmt.Errorf("io.Copy: %v", err)
// 	}
// 	fmt.Println(wc.Name)
// 	if err := wc.Close(); err != nil {
// 		return fmt.Errorf("Writer.Close: %v", err)
// 	}

// 	return nil
// }
