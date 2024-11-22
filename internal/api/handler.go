package api

import (
	"encoding/json"
	"net/http"

	"image_resampler/internal/config"
	"image_resampler/internal/processor"
	"image_resampler/internal/storage"
	"image_resampler/internal/validation"
)

type ImageHandler struct {
	Config *config.Config
}

func NewRouter(cfg *config.Config) http.Handler {
	router := http.NewServeMux()
	handler := &ImageHandler{
		Config: cfg,
	}

	router.HandleFunc("/process", handler.ProcessImage)

	return router
}

func (r *ImageHandler) ProcessImage(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var payload struct {
		Image string `json:"image"`
	}

	if err := json.NewDecoder(req.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}
	imageBytes, err := validation.ValidateImagePayload(payload.Image)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// imageBytes, err := base64.StdEncoding.DecodeString(payload.Image)
	// if err != nil {
	// 	http.Error(w, "Invalid base64 encoding", http.StatusBadRequest)
	// 	return
	// }
	var (
		cached         bool
		processingTime int64
	)
	if storage.CheckCache(imageBytes, int(r.Config.Width), int(r.Config.Height), r.Config.ResDir) {
		processedImg, procTime, isCached, err := processor.Process(imageBytes, r.Config.Width, r.Config.Height)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		processingTime = procTime
		cached = isCached
		storage.SaveImages(processedImg, imageBytes, r.Config.ResDir, r.Config.OrigDir, int(r.Config.Width), int(r.Config.Height))
	} else {
		cached = true
		processingTime = 0
	}

	response := map[string]interface{}{
		"time":   processingTime,
		"cached": cached,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
