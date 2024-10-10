package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/chetanji028/distributed-file-storage/internal/service"
	"github.com/google/uuid"
)

type FileHandler struct {
	Service service.FileService
}

func NewFileHandler(service service.FileService) *FileHandler {
	return &FileHandler{Service: service}
}

// UploadFileHandler handles the file upload
func (h *FileHandler) UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	fileData, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read file data", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	fileID, err := h.Service.UploadFile(fileData)
	if err != nil {
		http.Error(w, "Failed to upload file", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"file_id": fileID.String()}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetFileDataHandler retrieves metadata or parts information
func (h *FileHandler) GetFileDataHandler(w http.ResponseWriter, r *http.Request) {
	// Implement as needed. For simplicity, returning a placeholder.
	response := map[string]string{"message": "Get file data endpoint"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// DownloadFileHandler handles file download by ID
func (h *FileHandler) DownloadFileHandler(w http.ResponseWriter, r *http.Request) {
	fileIDStr := r.URL.Query().Get("id")
	if fileIDStr == "" {
		http.Error(w, "File ID is required", http.StatusBadRequest)
		return
	}

	fileID, err := uuid.Parse(fileIDStr)
	if err != nil {
		http.Error(w, "Invalid File ID", http.StatusBadRequest)
		return
	}

	fileData, err := h.Service.DownloadFile(fileID)
	if err != nil {
		http.Error(w, "Failed to download file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename="+fileID.String())
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fileData)
}
