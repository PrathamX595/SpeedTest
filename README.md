# Network Speed Test Web Service

This project provides a simple server implementation for conducting network speed tests. It measures download and upload speeds between the client and server, helping diagnose network performance issues.

## Features

- **Download Speed Test**: Sends a fixed amount of data to the client to measure download speed
- **Upload Speed Test**: Receives data from the client to measure upload speed
- **Results API**: Returns detailed information about test results

## Technical Details

The service is built using Go with the following components:
- Gorilla Mux for routing
- Clean architecture with controllers and routers
- CORS support for cross-origin requests
- Standardized JSON response format

## API Endpoints

### Download Test
```
GET /download
```
Sends 1MB of data to the client for download speed measurement.

### Download Results
```
GET /downloadResults
```
Returns the results of the most recent download test.

### Upload Test
```
POST /upload
```
Receives data from client and calculates upload speed.

## Response Format

All tests return a JSON response with the following structure:
```json
{
    "bytesReceived": 1024000,
    "megabytesReceived": 0.9765625,
    "durationSeconds": 0.5,
    "speedMbps": 15.625,
    "timestamp": "2023-04-01T12:00:00Z"
}
```

## Getting Started

1. Ensure Go 1.24.0 or later is installed
2. Clone the repository
3. Run the server:
     ```
     go run server.go
     ```
4. The server will start on port 5000

## Usage

The API can be used with any HTTP client. For web applications, you can make requests to the endpoints to measure network performance.