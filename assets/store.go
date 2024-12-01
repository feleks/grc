package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	assets := make(map[string]string)
	err := filepath.Walk("./assets", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			path2 := strings.Replace(path, "\\", "/", -1)
			if path2 == "assets/store.go" {
				return nil
			}

			content, err := os.ReadFile(path)
			if err != nil {
				return fmt.Errorf("failed to read file %s: %w", path, err)
			}

			b64Content := base64.StdEncoding.EncodeToString(content)
			if len(b64Content) == 0 {
				return fmt.Errorf("failed to base64 decode file %s: %w", path, err)
			}

			assets[path2] = b64Content
		}

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	goFileContent := "package main\n\nvar assetsContent = map[string]string{\n"

	for k, v := range assets {
		goFileContent += fmt.Sprintf(`    "%s": "%s"`, k, v) + ",\n"
	}

	goFileContent += "}\n"

	err = os.WriteFile("./assets_contents.go", []byte(goFileContent), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}
