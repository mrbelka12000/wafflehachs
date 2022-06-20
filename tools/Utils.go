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
		for i, line := range lines {
			if len(strings.TrimSpace(line)) > 0 && !strings.HasPrefix(line, "#") {
				arr := strings.Split(line, "=")
				if len(arr) != 2 {
					log.Fatalf("invalid format at line %v\n%v", i+1, line)
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
