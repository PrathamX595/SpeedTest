package main

import (
	"diagnosis/router"
	"net/http"
)

func main() {
	r := router.Router()
	http.ListenAndServe(":5000", r)
}
