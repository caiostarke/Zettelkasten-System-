package indexer

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Indexer indexes files in a directory and returns an HTML File with a list of all files in the directory
func IndexFiles(path string) (string, error) {
	var htmlContent string

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() || info.Name() == "style.css" {
			return nil
		}

		link := fmt.Sprintf(`<a href="%s">%s</a><br>`, strings.Replace(path, "storage/", "notes/", 1), info.Name())

		htmlContent += link

		return nil
	})

	return htmlContent, err
}
