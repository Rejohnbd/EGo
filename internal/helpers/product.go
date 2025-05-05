package helpers

import (
	"dakbazar/database"
	"dakbazar/internal/models"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

func GetStaticOption(key string, defaultValue *string) string {
	var option models.StaticOption
	err := database.DBMsql.Model(&models.StaticOption{}).
		Where("option_name = ?", key).
		First(&option).Error

	if err != nil {
		return *defaultValue
	}

	return *option.OptionValue
}

// func TaxOptionsSumRate(item *models.Product, countryID, stateID, cityID uint64) *models.Product {
// 	var (
// 		taxPercentage         float64
// 		taxPercentageForState float64
// 		taxPercentageForCity  float64
// 	)

// 	// Filtered slices
// 	var taxOptions, stateTaxOptions, cityTaxOptions []models.TaxClassOption

// 	for _, taxOption := range item.TaxOptions {
// 		if taxOption.CountryID != nil && *taxOption.CountryID == countryID {
// 			taxOptions = append(taxOptions, taxOption)
// 			taxPercentage += taxOption.Rate
// 		}
// 		if taxOption.StateID != nil && *taxOption.StateID == stateID {
// 			stateTaxOptions = append(stateTaxOptions, taxOption)
// 			taxPercentageForState += taxOption.Rate
// 		}
// 		if taxOption.CityID != nil && *taxOption.CityID == cityID {
// 			cityTaxOptions = append(cityTaxOptions, taxOption)
// 			taxPercentageForCity += taxOption.Rate
// 		}
// 	}

// 	// Default sum if no location provided
// 	if len(item.TaxOptions) > 0 && countryID == 0 {
// 		var total float64
// 		for _, t := range item.TaxOptions {
// 			total += t.Rate
// 		}
// 		item.TaxOptionsSumRate = total
// 		return item
// 	}

// 	// Country-level tax
// 	if countryID != 0 && len(taxOptions) > 0 {
// 		item.TaxOptionsSumRate = taxPercentage
// 	}

// 	// State-level tax
// 	if stateID != 0 && len(stateTaxOptions) > 0 {
// 		item.TaxOptionsSumRate = taxPercentageForState
// 	}

// 	// City-level tax
// 	if cityID != 0 && len(cityTaxOptions) > 0 {
// 		item.TaxOptionsSumRate = taxPercentageForCity
// 	}

// 	return item
// }

func AddonProductInstance() *gorm.DB {
	now := time.Now()

	query := database.DBMsql.Model(&models.Product{})

	query = query.Preload("Image").
		Preload("Badge").
		Preload("Uom").
		Preload("Uom.Unit").
		Preload("Vendor").
		Preload("Vendor.Status").
		Preload("TaxOptions").
		Preload("VendorAddress").
		Preload("ProductCategory.Category").
		Preload("ProductSubCategory.SubCategory").
		Preload("ProductChildCategory.ChildCategory").
		Preload("Inventory").
		Preload("Ratings").          //count need
		Preload("InventoryDetails"). //count need
		Preload("OrderItemsCount").  //count need
		Preload("CampaignProduct", func(db *gorm.DB) *gorm.DB {
			return db.Where("start_date <= ? AND end_date >= ?", now, now)
		})

	// if GetStaticOption("vendor_enable", "on") != "on" {
	// 	query = query.Where("vendor_id IS NULL")
	// }

	// query = query.Preload("TaxClass.Options")

	return query

	// $products->withAvg('ratings', 'rating');
	// $products->withSum("taxOptions", "rate")->when(get_static_option('vendor_enable', 'on') != 'on', function ($query) {
	//     $query->whereNull("vendor_id");
	// });
}

func ProductOrderItemQuery(products *gorm.DB, settings map[string]interface{}) ([]models.Product, error) {
	// 1. Get order_by, order, items
	orderBy := "id"
	order := "asc"
	items := 0

	if val, ok := settings["order_by"]; ok {
		orderBy = fmt.Sprintf("%v", val)
	}

	if val, ok := settings["order"]; ok {
		order = fmt.Sprintf("%v", val)
	}

	if val, ok := settings["items"]; ok {
		if strVal := fmt.Sprintf("%v", val); strVal != "" {
			fmt.Sscanf(strVal, "%d", &items)
		}
	}

	// 2. Apply orderBy
	allProducts := products.Order(orderBy + " " + strings.ToLower(order))

	// 3. Apply limit (take) if items not empty
	if items > 0 {
		allProducts = allProducts.Limit(items)
	}

	// 4. Fetch products
	var productList []models.Product
	if err := allProducts.Find(&productList).Error; err != nil {
		return nil, err
	}

	// 5. Transform (apply tax_options_sum_rate)
	// for i := range productList {
	// 	product := &productList[i]

	// 	if product.VendorID != nil && GetStaticOption("calculate_tax_based_on", nil) == "vendor_shop_address" {
	// 		// Vendor Address available
	// 		var vendorAddress models.VendorAddress
	// 		if err := database.DBMsql.First(&vendorAddress, "vendor_id = ?", product.VendorID).Error; err == nil {
	// 			*product = TaxOptionsSumRate(*product, vendorAddress.CountryID, vendorAddress.StateID, vendorAddress.CityID)
	// 		}
	// 	} else if product.VendorID == nil && GetStaticOption("calculate_tax_based_on", nil) == "vendor_shop_address" {
	// 		// Admin Shop Manage address
	// 		var adminShop models.AdminShopManage
	// 		if err := database.DBMsql.Select("id", "country_id", "state_id", "city").First(&adminShop).Error; err == nil {
	// 			*product = TaxOptionsSumRate(*product, adminShop.CountryID, adminShop.StateID, adminShop.City)
	// 		}
	// 	}
	// }

	return productList, nil
}
