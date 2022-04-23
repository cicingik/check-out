package model

const (
	Discount = `discount`
	FreeItem = `free_item`
)

type (
	Promo struct {
		Model
		Sku              string  `gorm:"Column:sku" json:"sku"`
		PromoType        string  `gorm:"Column:promo_type" json:"promo_type"`
		MinimalPurchased int     `gorm:"Column:minimal_purchased" json:"minimal_purchased"`
		BonusProductSku  int     `gorm:"Column:bonus_product_sku" json:"bonus_product_sku"`
		Discount         float64 `gorm:"Column:discount" json:"discount"`
	}
)

func (t *Promo) TableName() string {
	return "ecommerce.promo"
}
