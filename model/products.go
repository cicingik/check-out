package model

type (
	Product struct {
		Model
		Sku      string  `gorm:"Column:sku" json:"sku"`
		Name     string  `gorm:"Column:name" json:"name"`
		Price    float64 `gorm:"Column:price" json:"price"`
		Quantity int     `gorm:"Column:quantity" json:"quantity"`
	}
)

func (t *Product) TableName() string {
	return "ecommerce.products"
}
