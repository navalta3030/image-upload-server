package routes

import (
	"encoding/json"
	"net/http"

	controllers "image_upload_server/controllers/ImageUpload"
)

// ImageUploadRoute - shall generate token
func ImageUploadRoute(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(10 * 1024 * 1024)
	// file, handler, err := r.FormFile("file")
	// get the data and information of user
	files := r.MultipartForm.File["file"]
	userForm := r.FormValue("meta_data")

	// Predict Image
	imageResponse := controllers.ImagePredictAPI(files)

	// Upload the image within digital ocean
	statusCode := controllers.ImageUploadDigitalOcean(files, userForm)

	// Insert to database
	controllers.InsertLinkToDatabase(userForm, imageResponse)

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(&imageResponse)
}
