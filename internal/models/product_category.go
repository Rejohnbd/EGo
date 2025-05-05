package models

type ProductCategory struct {
	ID         uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductID  uint64 `gorm:"not null;index" json:"product_id"`
	CategoryID uint64 `gorm:"not null;index" json:"category_id"`

	Category Category `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
}
