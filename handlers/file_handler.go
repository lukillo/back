package handlers

import (
	"bytes"
	"encoding/csv"
	"io"
	"net/http"
	"strings"
	"log"
	"email-marketing-api/utils"
)

// UploadAttachmentHandler handles binary file uploads for CSV, JPG, and PDF
func UploadAttachmentHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the filename from the query parameter
	filename := r.URL.Query().Get("filename")
	if filename == "" {
		http.Error(w, "Filename query parameter is required", http.StatusBadRequest)
		return
	}

	// Read the binary content from the request body
	var buf bytes.Buffer
	_, err := io.Copy(&buf, r.Body)
	if err != nil {
		http.Error(w, "Failed to read the file", http.StatusInternalServerError)
		return
	}

	// Handle file type based on the filename extension
	if strings.HasSuffix(strings.ToLower(filename), ".csv") {
		handleCSVUpload(w, &buf)
	} else if strings.HasSuffix(strings.ToLower(filename), ".jpg") || strings.HasSuffix(strings.ToLower(filename), ".pdf") {
		// handleAttachmentUpload(w, filename, buf.Bytes())
	} else {
		http.Error(w, "Only CSV, JPG, and PDF files are allowed", http.StatusBadRequest)
		return
	}
}

// handleCSVUpload processes the uploaded CSV file
func handleCSVUpload(w http.ResponseWriter, buf *bytes.Buffer) {
	reader := csv.NewReader(buf)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			log.Println("ERROR", err)
			break
		}
		if err != nil {
			log.Println("ERROR", err)
			http.Error(w, "Invalid CSV format", http.StatusBadRequest)
			return
		}
		log.Println("EMAIL", record[0])
		email := strings.TrimSpace(record[0])
		if email != "" {
			utils.AddEmail(email)
		}
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("CSV processed and emails added successfully"))
}

