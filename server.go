package main

import (
	"fmt"
	"mime"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

func index(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if path == "/" {
		assetsEndpoint(w, r, "/assets/index.html")
	} else if strings.HasPrefix(path, "/assets/") {
		assetsEndpoint(w, r, "")
	} else {
		http.NotFound(w, r)
	}

	//_, err := os.Stat("./index.html")
	//if err != nil {
	//	if !errors.Is(err, os.ErrNotExist) {
	//		http.Error(w, fmt.Sprintf("failed to check file existance: %s", err.Error()), http.StatusInternalServerError)
	//	}
	//} else {
	//	http.ServeFile(w, r, "./index.html")
	//	return
	//}
	//
	//indexHTML, err := base64.StdEncoding.DecodeString(indexHTMLInB64)
	//if err != nil {
	//	http.Error(w, fmt.Sprintf("failed to decode indexHTML: %s", err.Error()), http.StatusInternalServerError)
	//}
	//
	//w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//w.WriteHeader(http.StatusOK)
	//_, err = w.Write([]byte(indexHTML))
	//if err != nil {
	//	http.Error(w, fmt.Sprintf("failed to write response: %s", err.Error()), http.StatusInternalServerError)
	//}
}

func assetsEndpoint(w http.ResponseWriter, r *http.Request, path string) {
	if path == "" {
		path = r.URL.Path
	}
	path = path[1:]
	fileName := strings.TrimPrefix(path, "assets/")

	content, err := getAsset(path)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to get asset: %s", err.Error()), http.StatusNotFound)
		return
	}

	contentType := mime.TypeByExtension(filepath.Ext(fileName))
	if contentType == "" {
		contentType = http.DetectContentType(content)
	}
	if contentType == "" {
		contentType = "application/octet-stream"
	}
	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Length", strconv.Itoa(len(content)))
	w.Header().Set("Cache-Control", "no-cache")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(content)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to write response: %s", err.Error()), http.StatusInternalServerError)
		return
	}
}

func setupRoutes() {
	http.HandleFunc("/", index)
	http.HandleFunc("/ws", wsEndpoint)
}
