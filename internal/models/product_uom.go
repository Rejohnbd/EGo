package models

// ProductUOM represents a row in the product_uom table
type ProductUOM struct {
	ID        uint64  `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	ProductID uint64  `gorm:"index;not null;column:product_id" json:"product_id"`
	UnitID    uint64  `gorm:"index;not null;column:unit_id" json:"unit_id"`
	Quantity  float64 `gorm:"type:double(8,2);not null;column:quantity" json:"quantity"`

	Unit *Unit `gorm:"foreignKey:UnitID" json:"unit,omitempty"`
}

// TableName overrides the table name used by GORM
func (ProductUOM) TableName() string {
	return "product_uom"
}
