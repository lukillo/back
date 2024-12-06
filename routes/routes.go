package routes

import (
	"github.com/gorilla/mux"
	"email-marketing-api/handlers"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/upload", handlers.UploadAttachmentHandler).Methods("POST")
}