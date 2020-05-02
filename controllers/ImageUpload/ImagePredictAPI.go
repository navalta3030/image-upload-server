package controllers

import (
	"encoding/json"
	"image_upload_server/constants"
	"image_upload_server/models"
	"io"

	"bytes"
	"mime/multipart"
	"net/http"
)

// ImagePredictAPI - db connection
func ImagePredictAPI(files []*multipart.FileHeader) models.ImagePredictAPI {
	var data models.ImagePredictAPI
	var b bytes.Buffer

	// Create buffer
	w := multipart.NewWriter(&b)

	// copy files
	for _, file := range files {

		fileWriter, err := w.CreateFormFile("file", file.Filename)

		if err != nil {
			logError.Println(err)
		}

		// open and copy
		f, _ := file.Open()
		io.Copy(fileWriter, f)

	}
	w.Close()

	// Initalize request template
	request, err := http.NewRequest("POST", constants.SanicServer+"/predictImage", &b)
	if err != nil {
		logError.Println(err)
	}

	// set header and perform request
	request.Header.Set("Content-Type", w.FormDataContentType())
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		logError.Println(err)
	}
	defer response.Body.Close()

	// Parse request
	err1 := json.NewDecoder(response.Body).Decode(&data)
	if err1 != nil {
		logError.Println(err1)
	}

	return data
}
