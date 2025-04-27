package helpers

import (
	"dakbazar/database"
	"dakbazar/internal/models"
)

type ApiResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SafeString(val *string) string {
	if val != nil {
		return *val
	}
	return ""
}

// func uintToString(val *uint) string {
// 	if val != nil {
// 		return strconv.FormatUint(uint64(*val), 10)
// 	}
// 	return ""
// }

// func formatTime(t *time.Time) string {
// 	if t != nil {
// 		return t.Format("2006-01-02 15:04:05")
// 	}
// 	return ""
// }

func GetImageDetails(imageId string) models.MediaUpload {
	var assetsUrl = "assets/uploads/media-uploader/"
	var baseUrl = "http://127.0.0.1:3000/"

	var media models.MediaUpload

	database.DBMsql.First(&media, imageId)

	image := models.MediaUpload{
		Title:      media.Title,
		Path:       baseUrl + assetsUrl + media.Path,
		Alt:        media.Alt,
		Size:       media.Size,
		Dimensions: media.Dimensions,
		VendorID:   media.VendorID,
		UserID:     media.UserID,
		CreatedAt:  media.CreatedAt,
		UpdatedAt:  media.UpdatedAt,
	}

	return image
}
