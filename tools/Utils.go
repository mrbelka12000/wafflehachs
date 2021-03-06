package tools

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"syscall"

	uuid "github.com/satori/go.uuid"
)

func Loadenv(files ...string) {
	alseitov(files...)
}

// alseitov godotenv implementation
func alseitov(files ...string) {
	if files == nil {
		files = append(files, ".env")
	}

	for _, file := range files {
		bytes, err := ioutil.ReadFile(file)
		if err != nil {
			log.Printf("error parsing file: %v \n", err.Error())
			continue
		}
		lines := strings.Split(string(bytes), "\n")
		for _, line := range lines {
			if len(strings.TrimSpace(line)) > 0 && !strings.HasPrefix(line, "#") {
				arr := strings.Split(line, "=")
				if len(arr) > 2 {
					arr[1] += "=" + arr[2]
				}
				key, value := arr[0], arr[1]
				os.Setenv(key, value)
			}
		}
	}
}

func MakeJsonString(value interface{}) string {
	if value == nil {
		return "{}"
	}
	bf := bytes.NewBufferString("")
	e := json.NewEncoder(bf)
	e.SetEscapeHTML(false)
	e.Encode(value)
	return bf.String()
}

func GetPointerString(value string) *string {
	return &value
}

func GetRandomString() string {
	return uuid.NewV4().String()
}

func IsValidType(file multipart.File) (string, bool) {
	buff := make([]byte, 512)
	if _, err := file.Read(buff); err != nil {
		return "", false
	}
	filetype := http.DetectContentType(buff)
	if filetype != "image/jpeg" && filetype != "image/gif" && filetype != "image/png" {
		return "", false
	}
	return filetype, true
}

func GetStorageUrl(filename string) string {
	return fmt.Sprintf("%v/%v/%v/%v", os.Getenv("googlestorage"), os.Getenv("bucket"), os.Getenv("uploadPath"), filename)
}

func GetFileNameFromUrl(fileUrl string) string {
	parsedUrl := strings.Split(fileUrl, "/")
	return parsedUrl[len(parsedUrl)-1]
}

func CheckSignal(signal os.Signal) bool {
	switch signal {
	case os.Interrupt:
		return true
	case syscall.SIGINT:
		return true
	case syscall.SIGTERM:
		return true
	default:
		return false
	}
}
