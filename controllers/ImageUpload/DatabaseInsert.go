package controllers

import (
	"encoding/json"
	"fmt"
	"image_upload_server/constants"
	"image_upload_server/models"
	"image_upload_server/utils"
)

// InsertLinkToDatabase - db connection
func InsertLinkToDatabase(user string, predictions models.ImagePredictAPI) {
	var imageLink models.ImageLink
	var userModel models.User

	// parse incoming user json
	errParse := json.Unmarshal([]byte(user), &userModel)
	if errParse != nil {
		logError.Fatal(errParse)
	}

	// Create if not exist
	utils.UseDb().Where(&userModel).FirstOrCreate(&userModel)

	for _, prediction := range predictions.Data {

		imageLink = models.ImageLink{
			Link:                fmt.Sprintf("%s%s_%s", constants.ImageLinkS3, userModel.Email, prediction[0]),
			AssistantPrediction: prediction[1],
			UserID:              userModel.ID,
		}

		// Create if not exist
		utils.UseDb().Where(&imageLink).FirstOrCreate(&imageLink)
	}
}
