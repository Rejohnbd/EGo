package models

type ProductInventoryDetails struct {
	ID                 uint64   `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	ProductInventoryID uint64   `gorm:"not null;index;column:product_inventory_id" json:"product_inventory_id"`
	ProductID          uint64   `gorm:"not null;index;column:product_id" json:"product_id"`
	Color              *string  `gorm:"type:varchar(191);column:color" json:"color"`
	Size               *string  `gorm:"type:varchar(191);column:size" json:"size"`
	Hash               *string  `gorm:"type:text;column:hash" json:"hash"`
	AdditionalPrice    float64  `gorm:"type:double(8,2);default:0.00;column:additional_price" json:"additional_price"`
	AddCost            *float64 `gorm:"type:double(8,2);column:add_cost" json:"add_cost"`
	Image              *uint64  `gorm:"column:image" json:"image"`
	StockCount         int64    `gorm:"column:stock_count;default:0" json:"stock_count"`
	SoldCount          int64    `gorm:"column:sold_count;default:0" json:"sold_count"`
}

func (ProductInventoryDetails) TableName() string {
	return "product_inventory_details"
}
