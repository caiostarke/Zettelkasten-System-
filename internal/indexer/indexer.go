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

		if info.Name() == "styles" || strings.Contains(info.Name(), "css") {
			return nil
		}

		var link string

		if info.IsDir() && info.Name() != "storage" && info.Name() != "categories" && info.Name() != "tags" {
			link = fmt.Sprintf(`<div class="folder"><span>Folder</span><h2>%s</h2></div>`, info.Name())
		} else if !info.IsDir() {
			link = fmt.Sprintf(`<a href="%s">%s</a>`, strings.Replace(path, "storage/", "notes/", 1), strings.Split(info.Name(), ".")[0])
		}

		htmlContent += link

		return nil
	})

	return htmlContent, err
}
