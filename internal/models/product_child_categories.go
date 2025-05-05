package models

type ProductChildCategory struct {
	ID              uint64 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	ProductID       uint64 `gorm:"not null;index;column:product_id" json:"product_id"`
	ChildCategoryID uint64 `gorm:"not null;index;column:child_category_id" json:"child_category_id"`

	// Optional Relations
	Product       *Product       `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	ChildCategory *ChildCategory `gorm:"foreignKey:ChildCategoryID" json:"child_category,omitempty"`
}

func (ProductChildCategory) TableName() string {
	return "product_child_categories"
}
