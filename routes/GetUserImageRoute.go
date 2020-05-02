package routes

import (
	"encoding/json"
	"net/http"

	controllers "image_upload_server/controllers/GetUserImage"
	"image_upload_server/models"
)

// GetUserImageRoute - should return collection of images
func GetUserImageRoute(w http.ResponseWriter, r *http.Request) {
	var imageCollection models.ImageCollection

	// Get user images from database
	imageCollection = controllers.GetUserImage(r.Body)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&imageCollection)
}
