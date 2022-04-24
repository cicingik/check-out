package database

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type (
	Cart struct {
		Model
		Sku      string `gorm:"Column:sku" json:"sku"`
		Quantity int    `gorm:"Column:quantity" json:"quantity"`
		Status   int    `gorm:"Column:status" json:"status"`
		Discount float64
	}

	CartPagination struct {
		Cart []Cart `json:"cart"`
		Pg   struct {
			After     int `json:"after"`
			Before    int `json:"before"`
			TotalPage int `json:"total"`
			Page      int `json:"page"`
		} `json:"pagination"`
	}

	BundleCart struct {
		db *gorm.DB
		t  Cart
	}
)

func (t *Cart) TableName() string {
	return "ecommerce.cart"
}

func InitCart(ctx context.Context, g *gorm.DB) *BundleCart {
	return &BundleCart{
		db: g.WithContext(ctx),
		t:  Cart{},
	}
}

func (b *BundleCart) FindAll(data RequestPaginatorParam) (r CartPagination, err error) {

	var cart []Cart

	tx := b.db.
		Model(Promo{}).
		Select(`"ecommerce"."cart".*`)

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
	}, &cart)

	r.Cart = cart
	r.Pg.After = paging.NextPage
	r.Pg.Before = paging.PrevPage
	r.Pg.TotalPage = paging.TotalPage
	r.Pg.Page = data.Page

	return r, nil
}

func (b *BundleCart) FindCart(column string, value string) (Cart, error) {
	var err error
	var cart Cart

	query := fmt.Sprintf(`%s = ?`, column)
	err = b.db.Model(Cart{}).Where(query, value).First(&cart).Error
	if err != nil {
		return cart, err
	}

	return cart, err
}
