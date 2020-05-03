package controllers

import (
	"encoding/json"
	"io"

	"image_upload_server/models"
	"image_upload_server/utils"
)

// GetUserImage - db connection
func GetUserImage(requestBody io.ReadCloser) models.ImageCollection {
	var user models.User
	var imageCollection models.ImageCollection

	// parse json body
	err := json.NewDecoder(requestBody).Decode(&user)
	if err != nil {
		logError.Println(err)
	}

	logger.Println("Getting information image for " + user.Email)

	// query join
	dbErr := utils.UseDb().
		Table("image_links").
		Select("image_links.link, image_links.assistant_prediction, image_links.doctor_prediction").
		Joins("join users on users.id = image_links.user_id").
		Where("users.email = ?", user.Email).
		Scan(&imageCollection.Data)

	// error
	if dbErr.Error != nil {
		logError.Println(dbErr.Error)
		return models.ImageCollection{}
	}

	if dbErr.RecordNotFound() {
		return models.ImageCollection{}
	}

	return imageCollection
}
