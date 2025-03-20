package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type SpeedTestResult struct {
	BytesReceived     int     `json:"bytesReceived"`
	MegabytesReceived float64 `json:"megabytesReceived"`
	DurationSeconds   float64 `json:"durationSeconds"`
	SpeedMbps         float64 `json:"speedMbps"`
	Timestamp         string  `json:"timestamp"`
}

var result SpeedTestResult

func Download(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	chunkSize := 1024
	chunks := 1000
	totalBytes := 1024000

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", totalBytes))

	data := make([]byte, chunkSize)

	start := time.Now()

	for i := 0; i < chunks; i++ {
		_, err := w.Write(data)
		if err != nil {
			return
		}

		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}
	}

	duration := time.Since(start).Seconds()
	megabytesReceived := float64(totalBytes) / (1024 * 1024)
	speedMbps := (megabytesReceived * 8) / duration

	result = SpeedTestResult{
		BytesReceived:     totalBytes,
		MegabytesReceived: megabytesReceived,
		DurationSeconds:   duration,
		SpeedMbps:         speedMbps,
		Timestamp:         time.Now().Format(time.RFC3339),
	}
}

func DownloadResults(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func Upload(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	start := time.Now()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading request body: %v", err), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	bytesReceived := len(body)
	duration := time.Since(start).Seconds()
	mbReceived := float64(bytesReceived) / (1024 * 1024)
	mbps := (mbReceived * 8) / duration

	result := SpeedTestResult{
		BytesReceived:     bytesReceived,
		MegabytesReceived: mbReceived,
		DurationSeconds:   duration,
		SpeedMbps:         mbps,
		Timestamp:         time.Now().Format(time.RFC3339),
	}

	json.NewEncoder(w).Encode(result)
}
