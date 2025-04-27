package controllers

import (
	"dakbazar/database"
	"dakbazar/internal/helpers"
	"dakbazar/internal/models"
	"strings"

	"github.com/elliotchance/phpserialize"
	"github.com/gofiber/fiber/v2"
)

func Sliders(c *fiber.Ctx) error {
	var allWidgets []models.PageBuilder

	allWidgetsResult := database.DBMsql.
		Where("addon_page_type = ?", "dynamic_page").
		Where("addon_page_id = ?", 3).
		Order("id DESC").
		Find(&allWidgets)

	if allWidgetsResult.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": allWidgetsResult.Error.Error(),
		})
	}

	type Slider struct {
		Subtitle string `json:"subtitle"`
		Title    string `json:"title"`
		BtnText  string `json:"btn_text"`
		BtnURL   string `json:"btn_url"`
		Image    struct {
			Path  string `json:"path"`
			Alt   string `json:"alt"`
			Title string `json:"title"`
		} `json:"image"`
	}

	var finalSliders []Slider

	for _, widget := range allWidgets {
		if widget.AddonSettings == nil {
			continue
		}

		var settings map[interface{}]interface{}
		err := phpserialize.Unmarshal([]byte(*widget.AddonSettings), &settings)
		if err != nil {
			continue
		}

		headerSlider, ok := settings["header_slider"].(map[interface{}]interface{})
		if !ok {
			continue
		}

		subtitles, _ := headerSlider["subtitle_"].([]interface{})
		titles, _ := headerSlider["title_"].([]interface{})
		btnTexts, _ := headerSlider["btn_text_"].([]interface{})
		btnUrls, _ := headerSlider["btn_url_"].([]interface{})
		images, _ := headerSlider["image_"].([]interface{})

		for i := 0; i < len(subtitles); i++ {
			slider := Slider{
				Subtitle: getStringFromInterface(subtitles, i),
				Title:    cleanTitle(getStringFromInterface(titles, i)),
				BtnText:  getStringFromInterface(btnTexts, i),
				BtnURL:   getStringFromInterface(btnUrls, i),
			}
			imgID := getStringFromInterface(images, i)

			if imgID != "" {
				imageInfo := helpers.GetImageDetails(imgID)
				slider.Image.Path = imageInfo.Path
				slider.Image.Alt = helpers.SafeString(imageInfo.Alt)
				slider.Image.Title = imageInfo.Title
			}

			finalSliders = append(finalSliders, slider)
		}
	}

	return c.JSON(helpers.ApiResponse{
		Status:  true,
		Message: "Sliders Data",
		Data:    finalSliders,
	})
}

func getStringFromInterface(arr []interface{}, index int) string {
	if index < len(arr) {
		if val, ok := arr[index].(string); ok {
			return val
		}
	}
	return ""
}

func cleanTitle(s string) string {
	// Remove [cl] and [/cl] from title
	s = strings.ReplaceAll(s, "[cl]", "")
	s = strings.ReplaceAll(s, "[/cl]", "")
	return s
}
