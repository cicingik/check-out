package model

type (
	Cart struct {
		Model
		Sku      string `gorm:"Column:sku" json:"sku"`
		Quantity int    `gorm:"Column:quantity" json:"quantity"`
		Status   int    `gorm:"Column:status" json:"status"`
	}
)

func (t *Cart) TableName() string {
	return "ecommerce.cart"
}
