package models

type SubOrderItem struct {
	ID         uint64   `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	SubOrderID uint64   `gorm:"not null;index;column:sub_order_id" json:"sub_order_id"`
	OrderID    uint64   `gorm:"not null;index;column:order_id" json:"order_id"`
	ProductID  *uint64  `gorm:"index;column:product_id" json:"product_id"`
	VariantID  *uint64  `gorm:"index;column:variant_id" json:"variant_id"`
	Quantity   uint64   `gorm:"not null;index;column:quantity" json:"quantity"`
	Price      float64  `gorm:"type:decimal(8,2);not null;index;column:price" json:"price"`
	SalePrice  float64  `gorm:"type:decimal(8,2);not null;index;column:sale_price" json:"sale_price"`
	TaxAmount  float64  `gorm:"type:decimal(8,2);default:0.00;column:tax_amount" json:"tax_amount"`
	TaxType    *string  `gorm:"type:varchar(191);column:tax_type" json:"tax_type"`
	Weight     *float64 `gorm:"type:double;column:weight" json:"weight"`
}

// TableName sets the correct table name
func (SubOrderItem) TableName() string {
	return "sub_order_items"
}
