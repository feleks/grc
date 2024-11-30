package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"os"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	_, err := os.Stat("grc/index.html")
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			http.Error(w, fmt.Sprintf("failed to check file existance: %s", err.Error()), http.StatusInternalServerError)
		}
	} else {
		http.ServeFile(w, r, "grc/index.html")
		return
	}

	indexHTML, err := base64.StdEncoding.DecodeString(indexHTMLInB64)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to decode indexHTML: %s", err.Error()), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(indexHTML))
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to write response: %s", err.Error()), http.StatusInternalServerError)
	}
}

func toBase64(w http.ResponseWriter, r *http.Request) {
	f, err := os.ReadFile("grc/index.html")
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to read file: %s", err.Error()), http.StatusInternalServerError)
	}

	fStr := base64.StdEncoding.EncodeToString(f)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(fStr))
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to write base64 file: %s", err.Error()), http.StatusInternalServerError)
	}
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
	http.HandleFunc("/2b64", toBase64)
}
