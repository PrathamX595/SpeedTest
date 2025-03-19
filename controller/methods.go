package controller

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func Download(w http.ResponseWriter, r *http.Request) {

	chunkSize := 1024
    chunks := 10000
    totalBytes := chunkSize * chunks

	w.Header().Set("Content-Length", fmt.Sprintf("%d", totalBytes))
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Content-Type", "application/octet-stream")


	data := make([]byte, chunkSize)
    for i := 0; i < chunks; i++ {
        w.Write(data)
    }
}

func Upload(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	r.ParseMultipartForm(10 << 20)
	file, _, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err1 := io.Copy(io.Discard, file)
	if err1 != nil {
		panic(err1)
	}
	duration := time.Since(start).Seconds()
	fmt.Fprintf(w, "Upload completed in %.2f seconds", duration)
}