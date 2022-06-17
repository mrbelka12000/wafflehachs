package tools

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
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
			log.Fatalf("error parsing file: %v", err.Error())
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
