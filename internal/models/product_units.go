package models

type ProductUnit struct {
	ID   uint64 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	Name string `gorm:"type:varchar(191);not null;index;column:name" json:"name"`
}

func (ProductUnit) TableName() string {
	return "product_units"
}
