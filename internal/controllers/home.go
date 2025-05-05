package controllers

import (
	"dakbazar/database"
	"dakbazar/internal/helpers"
	"dakbazar/internal/models"
	"fmt"
	"log"
	"strings"

	"github.com/elliotchance/phpserialize"
	"github.com/gofiber/fiber/v2"
)

func Sliders(c *fiber.Ctx) error {
	// Response Structure for Sliders
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

	var allWidgets []models.PageBuilder
	err := database.DBMsql.
		Where("addon_page_type = ?", "dynamic_page").
		Where("addon_page_id = ?", 3).
		Order("id DESC").
		Find(&allWidgets).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	var finalSliders []Slider

	for _, widget := range allWidgets {
		if widget.AddonSettings == nil {
			continue
		}

		var settings map[any]any
		if err := phpserialize.Unmarshal([]byte(*widget.AddonSettings), &settings); err != nil {
			continue
		}

		headerSlider, ok := settings["header_slider"].(map[any]any)
		if !ok {
			continue
		}

		subtitles, _ := headerSlider["subtitle_"].([]any)
		titles, _ := headerSlider["title_"].([]any)
		btnTexts, _ := headerSlider["btn_text_"].([]any)
		btnUrls, _ := headerSlider["btn_url_"].([]any)
		images, _ := headerSlider["image_"].([]any)

		for i, subtitle := range subtitles {
			var slider Slider

			if s, ok := subtitle.(string); ok {
				slider.Subtitle = s
			}
			if i < len(titles) {
				if t, ok := titles[i].(string); ok {
					slider.Title = cleanTitle(t)
				}
			}
			if i < len(btnTexts) {
				if b, ok := btnTexts[i].(string); ok {
					slider.BtnText = b
				}
			}
			if i < len(btnUrls) {
				if u, ok := btnUrls[i].(string); ok {
					slider.BtnURL = u
				}
			}
			if i < len(images) {
				if imgID, ok := images[i].(string); ok && imgID != "" {
					imageInfo := helpers.GetImageDetails(imgID)
					slider.Image.Path = imageInfo.Path
					slider.Image.Alt = helpers.SafeString(imageInfo.Alt)
					slider.Image.Title = imageInfo.Title
				}
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

func cleanTitle(s string) string {
	// Remove [cl] and [/cl] from title
	s = strings.ReplaceAll(s, "[cl]", "")
	s = strings.ReplaceAll(s, "[/cl]", "")
	return s
}

func FeaturedItems(c *fiber.Ctx) error {

	var pageBuilder models.PageBuilder

	pageBuilderResult := database.DBMsql.First(&pageBuilder, 179)

	if pageBuilderResult.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": pageBuilderResult.Error.Error(),
		})
	}

	// fmt.Println(*pageBuilder.AddonSettings)
	var addonSettings map[any]any
	err := phpserialize.Unmarshal([]byte(*pageBuilder.AddonSettings), &addonSettings)

	if err != nil {
		return c.Status(500).JSON(helpers.ApiResponse{
			Status:  false,
			Message: "Failed to unserialize AddonSettings",
			Data:    err.Error(),
		})
	}

	cleanAddonSettings := helpers.ConvertMapInterfaceToMapString(addonSettings)
	productsRaw, ok := cleanAddonSettings["product"]

	var productIds []any
	if ok {
		if castedProducts, ok := productsRaw.(map[string]any); ok {
			for _, v := range castedProducts {
				productIds = append(productIds, v)
			}
		}
	}

	productsQuery := helpers.AddonProductInstance()

	log.Printf("productIds: %v", productsQuery)

	if len(productIds) > 0 {
		productsQuery = productsQuery.Where("products.id IN ?", productIds)
	} else {
		productsQuery = productsQuery.Limit(10)
	}

	// --- Now call ProductOrderItemQuery ---
	finalProducts, err := helpers.ProductOrderItemQuery(productsQuery, cleanAddonSettings)

	if err != nil {
		return c.Status(500).JSON(helpers.ApiResponse{
			Status:  false,
			Message: "Failed to load products",
			Data:    err.Error(),
		})
	}

	type FeaturedItem struct {
		ID        uint     `json:"id"`
		Name      string   `json:"name"`
		Slug      string   `json:"slug"`
		Summary   *string  `json:"summary"`
		Price     *float64 `json:"price"`
		SalePrice *float64 `json:"sale_price"`
		Image     string   `json:"image"`
	}

	var featuredItems []FeaturedItem

	for _, product := range finalProducts {
		item := FeaturedItem{
			ID:        product.ID,
			Name:      product.Name,
			Slug:      product.Slug,
			Summary:   product.Summary,
			Price:     product.Price,     // *float64
			SalePrice: product.SalePrice, // *float64
			Image:     "",                // default empty
		}

		// Get image path if image exists
		if product.Image.Path != "" {
			item.Image = product.Image.Path
		}

		featuredItems = append(featuredItems, item)
	}

	return c.JSON(helpers.ApiResponse{
		Status:  true,
		Message: "Featured Items Data",
		Data:    featuredItems,
	})
}

func TodayDeals(c *fiber.Ctx) error {
	var pageBuilder models.PageBuilder

	pageBuilderResult := database.DBMsql.First(&pageBuilder, 180)

	if pageBuilderResult.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": pageBuilderResult.Error.Error(),
		})
	}

	// fmt.Println(*pageBuilder.AddonSettings)
	var addonSettings map[any]any
	err := phpserialize.Unmarshal([]byte(*pageBuilder.AddonSettings), &addonSettings)

	if err != nil {
		return c.Status(500).JSON(helpers.ApiResponse{
			Status:  false,
			Message: "Failed to unserialize AddonSettings",
			Data:    err.Error(),
		})
	}

	cleanAddonSettings := helpers.ConvertMapInterfaceToMapString(addonSettings)
	productsRaw, ok := cleanAddonSettings["product"]

	var productIds []any
	if ok {
		if castedProducts, ok := productsRaw.(map[string]any); ok {
			for _, v := range castedProducts {
				productIds = append(productIds, v)
			}
		}
	}

	productsQuery := helpers.AddonProductInstance()

	log.Printf("productIds: %v", productsQuery)

	if len(productIds) > 0 {
		productsQuery = productsQuery.Where("products.id IN ?", productIds)
	} else {
		productsQuery = productsQuery.Limit(10)
	}

	// --- Now call ProductOrderItemQuery ---
	finalProducts, err := helpers.ProductOrderItemQuery(productsQuery, cleanAddonSettings)

	if err != nil {
		return c.Status(500).JSON(helpers.ApiResponse{
			Status:  false,
			Message: "Failed to load products",
			Data:    err.Error(),
		})
	}

	type FeaturedItem struct {
		ID        uint     `json:"id"`
		Name      string   `json:"name"`
		Slug      string   `json:"slug"`
		Summary   *string  `json:"summary"`
		Price     *float64 `json:"price"`
		SalePrice *float64 `json:"sale_price"`
		Image     string   `json:"image"`
	}

	var featuredItems []FeaturedItem

	for _, product := range finalProducts {
		item := FeaturedItem{
			ID:        product.ID,
			Name:      product.Name,
			Slug:      product.Slug,
			Summary:   product.Summary,
			Price:     product.Price,     // *float64
			SalePrice: product.SalePrice, // *float64
			Image:     "",                // default empty
		}

		// Get image path if image exists
		if product.Image.Path != "" {
			item.Image = product.Image.Path
		}

		featuredItems = append(featuredItems, item)
	}

	return c.JSON(helpers.ApiResponse{
		Status:  true,
		Message: "Today Deals Data",
		Data:    featuredItems,
	})
}

func HotItems(c *fiber.Ctx) error {
	var pageBuilder models.PageBuilder

	pageBuilderResult := database.DBMsql.First(&pageBuilder, 182)

	if pageBuilderResult.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": pageBuilderResult.Error.Error(),
		})
	}

	// fmt.Println(*pageBuilder.AddonSettings)
	var addonSettings map[any]any
	err := phpserialize.Unmarshal([]byte(*pageBuilder.AddonSettings), &addonSettings)

	if err != nil {
		return c.Status(500).JSON(helpers.ApiResponse{
			Status:  false,
			Message: "Failed to unserialize AddonSettings",
			Data:    err.Error(),
		})
	}

	cleanAddonSettings := helpers.ConvertMapInterfaceToMapString(addonSettings)
	// fmt.Println(cleanAddonSettings)
	productsRaw, ok := cleanAddonSettings["products"]
	fmt.Println(cleanAddonSettings["products"])

	var productIds []any
	if ok {
		if castedProducts, ok := productsRaw.([]any); ok {
			productIds = append(productIds, castedProducts...)
		}
	}

	fmt.Println(productIds)

	productsQuery := helpers.AddonProductInstance()

	// log.Printf("productIds: %v", productsQuery)

	if len(productIds) > 0 {
		productsQuery = productsQuery.Where("products.id IN ?", productIds)
	} else {
		productsQuery = productsQuery.Limit(10)
	}

	// --- Now call ProductOrderItemQuery ---
	finalProducts, err := helpers.ProductOrderItemQuery(productsQuery, cleanAddonSettings)

	if err != nil {
		return c.Status(500).JSON(helpers.ApiResponse{
			Status:  false,
			Message: "Failed to load products",
			Data:    err.Error(),
		})
	}

	type FeaturedItem struct {
		ID        uint     `json:"id"`
		Name      string   `json:"name"`
		Slug      string   `json:"slug"`
		Summary   *string  `json:"summary"`
		Price     *float64 `json:"price"`
		SalePrice *float64 `json:"sale_price"`
		Image     string   `json:"image"`
	}

	type Category struct {
		ID          uint    `json:"id"`
		Name        string  `json:"name"`
		Slug        string  `json:"slug"`
		Description *string `json:"description"`
		ImageId     string  `json:"image_id,omitempty"`
		CreatedAt   string  `json:"created_at"`
		UpdatedAt   string  `json:"updated_at"`
		DeletedAt   string  `json:"deleted_at,omitempty"`
		StatusId    int16   `json:"status_id"`
	}

	var (
		featuredItems []FeaturedItem
		categories    []Category
		categoryMap   = make(map[uint]bool) // to check uniqueness
	)

	for _, product := range finalProducts {
		item := FeaturedItem{
			ID:        product.ID,
			Name:      product.Name,
			Slug:      product.Slug,
			Summary:   product.Summary,
			Price:     product.Price,
			SalePrice: product.SalePrice,
			Image:     "",
		}
		fmt.Println(product.ProductCategory)
		// Get image path if image exists
		if product.Image.Path != "" {
			item.Image = product.Image.Path
		}
		fmt.Printf("%v\n", product.ProductCategory.Category.Name)

		if product.ProductCategory.Category.ID != 0 {
			catID := uint(product.ProductCategory.Category.ID)
			if !categoryMap[catID] {
				categories = append(categories, Category{
					ID:          catID,
					Name:        product.ProductCategory.Category.Name,
					Slug:        product.ProductCategory.Category.Slug,
					Description: product.ProductCategory.Category.Description,
					// ImageId:     product.ProductCategory.Category.ImageID,
					CreatedAt: product.ProductCategory.Category.CreatedAt.String(),
					UpdatedAt: product.ProductCategory.Category.UpdatedAt.String(),
					// DeletedAt:   product.ProductCategory.Category.DeletedAt.String(),
					// StatusId:    product.ProductCategory.Category.StatusID,
				})
				categoryMap[catID] = true
			}
		}

		featuredItems = append(featuredItems, item)
	}

	return c.JSON(helpers.ApiResponse{
		Status:  true,
		Message: "Hot Items Data",
		Data: fiber.Map{
			"products":   featuredItems,
			"categories": categories,
		},
	})
}
