//
// Golang Workshop 2024
//

package main

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed web/*
var content embed.FS

func main() {
	http.Handle("/", http.FileServer(getFileSystem()))
	http.ListenAndServe(":8080", nil)
}

func getFileSystem() http.FileSystem {
	fsys, err := fs.Sub(content, "web")
	if err != nil {
		panic(err)
	}
	return http.FS(fsys)
}
