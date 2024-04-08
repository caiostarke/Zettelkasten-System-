package main

import (
	"Zettelkasten/internal/indexer"
	"fmt"
	"net/http"
	"os"

	"github.com/russross/blackfriday/v2"
)

const STORAGE_DIR = "./storage"

func main() {
	http.HandleFunc("/notes/", func(w http.ResponseWriter, r *http.Request) {
		filePath := r.URL.Path[len("/notes/"):]

		data, err := os.ReadFile(STORAGE_DIR + "/" + filePath)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		htmlContent := mdToHTML(data)

		w.Header().Set("Content-Type", "text/html")

		fmt.Fprintf(w, `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Markdown File</title>
			<link rel="stylesheet" href="/styles/style.css">
		</head>
		<body>
			%s <!-- Markdown content converted to HTML -->
		</body>
		</html>
	`, htmlContent)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		htmlContent, err := indexer.IndexFiles(STORAGE_DIR)
		if err != nil {
			http.Error(w, "Error indexing files", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html")

		fmt.Fprintf(w, `
		<!DOCTYPE html>
		<html>
		<head>
			<title>File Index</title>
			<link rel="stylesheet" href="/styles/style.css">
		</head>
		<body>
			<h1>File Index</h1>
			%s <!-- HTML content with links to files -->
		</body>
		</html>
	`, htmlContent)

	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func mdToHTML(file []byte) []byte {
	htmlContent := blackfriday.Run(file)

	return htmlContent
}
