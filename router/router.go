package router

import (
	"diagnosis/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/download", controller.Download)
	r.HandleFunc("/upload", controller.Upload)
	return r
}