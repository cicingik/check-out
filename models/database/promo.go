package database

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

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
		BonusProductSku  string  `gorm:"Column:bonus_product_sku" json:"bonus_product_sku"`
		Discount         float64 `gorm:"Column:discount" json:"discount"`
		IsActive         bool    `gorm:"Column:is_active" json:"is_active"`
	}

	PromoPagination struct {
		Promo []Promo `json:"promo"`
		Pg    struct {
			After     int `json:"after"`
			Before    int `json:"before"`
			TotalPage int `json:"total"`
			Page      int `json:"page"`
		} `json:"pagination"`
	}

	BundlePromo struct {
		db *gorm.DB
		t  Promo
	}
)

func (t *Promo) TableName() string {
	return "ecommerce.promo"
}

func InitPromo(ctx context.Context, g *gorm.DB) *BundlePromo {
	return &BundlePromo{
		db: g.WithContext(ctx),
		t:  Promo{},
	}
}

func (b *BundlePromo) Create(data Promo) (Promo, error) {
	var err error

	err = b.db.Create(&data).Error
	if err != nil {
		return Promo{}, err
	}

	return data, err
}

func (b *BundlePromo) FindAllWithPagination(data RequestPaginatorParam) (r PromoPagination, err error) {

	var promo []Promo

	tx := b.db.
		Select(`"ecommerce"."promo".*`).
		Where(`is_active = true`)

	err = tx.Error
	if err != nil {
		return r, err
	}

	paging := Paging(&PaginatorParam{
		DB:      tx,
		Page:    data.Page,
		Limit:   data.Limit,
		OrderBy: []string{`created_at desc`},
		ShowSQL: false,
	}, &promo)

	r.Promo = promo
	r.Pg.After = paging.NextPage
	r.Pg.Before = paging.PrevPage
	r.Pg.TotalPage = paging.TotalPage
	r.Pg.Page = data.Page

	return r, nil
}

func (b *BundlePromo) FindPromo(column string, value string) ([]Promo, error) {
	var err error
	var promo []Promo

	query := fmt.Sprintf(`%s = ?`, column)
	err = b.db.Model(Promo{}).Where(query, value).Find(&promo).Error
	if err != nil {
		return promo, err
	}

	return promo, err
}

func (b *BundlePromo) FindAll() (results []Promo, err error) {

	err = b.db.Find(&results).
		Where(`is_active = true`).
		Error

	if err != nil {
		return results, err
	}

	return results, nil
}

func (b *BundlePromo) Update(tx *gorm.DB, data Promo) (results Promo, err error) {
	err = tx.Model(&results).
		First(&results, data.ID).
		Updates(map[string]interface{}{
			"sku":               data.Sku,
			"promo_type":        data.PromoType,
			"minimal_purchased": data.MinimalPurchased,
			"bonus_product_sku": data.BonusProductSku,
			"discount":          data.Discount,
			"is_active":         data.IsActive,
		}).
		Error

	if err != nil {
		tx.Rollback()
		return results, err
	}

	return results, err
}
