package controller

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func Download(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Length", "10000000")
	w.Header().Set("Content-Type", "application/octet-stream")
	data := make([]byte, 1024)
	for i := 0; i < 10000; i++ {
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