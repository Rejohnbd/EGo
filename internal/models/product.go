package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID                  uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Name                string         `gorm:"type:varchar(191);not null;index" json:"name"`
	Slug                string         `gorm:"type:varchar(191);not null;index" json:"slug"`
	Summary             *string        `gorm:"type:text" json:"summary"`
	Title               *string        `gorm:"type:varchar(191);index" json:"title"`
	Description         *string        `gorm:"type:longtext" json:"description"`
	ImageID             *string        `gorm:"type:varchar(191)" json:"image_id"`
	Price               *float64       `gorm:"type:double" json:"price"`
	SalePrice           *float64       `gorm:"type:double" json:"sale_price"`
	Cost                *float64       `gorm:"type:double" json:"cost"`
	AdminID             *uint          `gorm:"column:admin_id;index" json:"admin_id"`
	VendorID            *uint          `gorm:"column:vendor_id;index" json:"vendor_id"`
	SupplierID          *uint          `gorm:"column:supplier_id;index" json:"supplier_id"`
	BadgeID             *uint          `gorm:"column:badge_id;index" json:"badge_id"`
	BrandID             *uint          `gorm:"column:brand_id;index" json:"brand_id"`
	TaxClassID          *uint          `gorm:"column:tax_class_id;index" json:"tax_class_id"`
	StatusID            uint           `gorm:"column:status_id;index;default:2" json:"status_id"`
	ProductType         uint           `gorm:"column:product_type;index;default:1" json:"product_type"`
	SoldCount           *int           `gorm:"type:int" json:"sold_count"`
	MinPurchase         *int           `gorm:"type:int" json:"min_purchase"`
	MaxPurchase         *int           `gorm:"type:int" json:"max_purchase"`
	IsRefundable        int16          `gorm:"type:smallint,column:is_refundable;index" json:"is_refundable"`
	IsInHouse           int16          `gorm:"type:smallint,column:is_in_house;index;default:1" json:"is_in_house"`
	IsInventoryWarnAble *int16         `gorm:"type:smallint,column:is_inventory_warn_able;index" json:"is_inventory_warn_able"`
	IsTaxable           int16          `gorm:"type:smallint,column:is_taxable;default:0" json:"is_taxable"`
	Weight              *float64       `gorm:"type:double" json:"weight"`
	CreatedAt           *time.Time     `gorm:"column:created_at" json:"created_at"`
	UpdatedAt           *time.Time     `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt           gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`

	Image                *MediaUpload              `gorm:"foreignKey:ImageID" json:"image"`
	Badge                *Badge                    `gorm:"foreignKey:BadgeID" json:"badge"`
	Uom                  *ProductUOM               `gorm:"foreignKey:ProductID;references:ID" json:"uom"`
	Vendor               *Vendor                   `gorm:"foreignKey:VendorID" json:"vendor"`
	TaxOptions           []TaxClassOption          `gorm:"foreignKey:ClassID;references:TaxClassID" json:"tax_options"`
	VendorAddress        *VendorAddress            `gorm:"foreignKey:VendorID;references:VendorID" json:"vendor_address"`
	CampaignSoldProduct  *CampaignSoldProduct      `gorm:"foreignKey:ProductID" json:"campaign_sold_product"`
	ProductCategory      *ProductCategory          `gorm:"foreignKey:ProductID" json:"category"`
	ProductSubCategory   *ProductSubCategory       `gorm:"foreignKey:ProductID" json:"sub_category"`
	ProductChildCategory *ProductChildCategory     `gorm:"foreignKey:ProductID" json:"child_category"`
	CampaignProduct      *CampaignProduct          `gorm:"foreignKey:ProductID" json:"campaign_product"`
	Inventory            *ProductInventory         `gorm:"foreignKey:ProductID" json:"inventory"`
	Ratings              []ProductRatings          `gorm:"foreignKey:ProductID;references:ID" json:"ratings"`
	InventoryDetails     []ProductInventoryDetails `gorm:"foreignKey:ProductID;references:ID" json:"inventory_details"`
	OrderItemsCount      []SubOrderItem            `gorm:"foreignKey:ProductID;references:ID" json:"order_items_count"`
	TaxClass             *TaxClass                 `gorm:"foreignKey:TaxClassID" json:"tax_class"`
}
