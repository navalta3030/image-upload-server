package models

// ImageCollectionQuery - database
type ImageCollectionQuery struct {
	Link                string `json:"link"`
	AssistantPrediction string `json:"assistant_prediction"`
	DoctorPrediction    string `json:"doctor_prediction"`
}

// ImageCollection - api response
type ImageCollection struct {
	Data []ImageCollectionQuery `json:"data"`
}
