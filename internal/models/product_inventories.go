package models

type ProductInventory struct {
	ID         uint64 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	ProductID  uint64 `gorm:"not null;index;column:product_id" json:"product_id"`
	SKU        string `gorm:"type:varchar(191);not null;index;column:sku" json:"sku"`
	StockCount *int   `gorm:"column:stock_count" json:"stock_count,omitempty"`
	SoldCount  *int   `gorm:"column:sold_count" json:"sold_count"`

	// Optional relationship
	Product *Product `gorm:"foreignKey:ProductID" json:"product,omitempty"`
}

// TableName sets the correct table name
func (ProductInventory) TableName() string {
	return "product_inventories"
}
