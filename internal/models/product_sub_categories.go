package models

type ProductSubCategory struct {
	ID            uint64 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	ProductID     uint64 `gorm:"not null;index;column:product_id" json:"product_id"`
	SubCategoryID uint64 `gorm:"not null;index;column:sub_category_id" json:"sub_category_id"`

	Product     *Product  `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	SubCategory *Category `gorm:"foreignKey:SubCategoryID" json:"sub_category"`
}

func (ProductSubCategory) TableName() string {
	return "product_sub_categories"
}
