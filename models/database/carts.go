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
		ClientId int    `gorm:"Column:client_id" json:"client_id"`
		Quantity int    `gorm:"Column:quantity" json:"quantity"`
		IsActive bool   `gorm:"Column:is_active" json:"is_active"`
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

	CheckOutItem struct {
		ClientId int    `json:"client_id"`
		Contents []Cart `json:"contents"`
	}

	BundleCart struct {
		db *gorm.DB
		t  Cart
	}
)

func (t *Cart) TableName() string {
	return "ecommerce.carts"
}

func InitCart(ctx context.Context, g *gorm.DB) *BundleCart {
	return &BundleCart{
		db: g.WithContext(ctx),
		t:  Cart{},
	}
}

func (b *BundleCart) Create(data Cart) (Cart, error) {
	var err error

	err = b.db.Create(&data).Error
	if err != nil {
		return Cart{}, err
	}

	return data, err
}

func (b *BundleCart) FindAll(data RequestPaginatorParam) (r CartPagination, err error) {

	var cart []Cart

	tx := b.db.
		Model(Promo{}).
		Where(`is_active = true`).
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
	err = b.db.Model(Cart{}).Where(query, value).Where(`is_active = true`).First(&cart).Error
	if err != nil {
		return cart, err
	}

	return cart, err
}

func (b *BundleCart) FindCartById(id int) (Cart, error) {
	var err error
	var cart Cart

	err = b.db.Model(Cart{}).
		Where(`id = ?`, id).
		Where(`is_active = true`).
		First(&cart).
		Error
	if err != nil {
		return cart, err
	}

	return cart, err
}

func (b *BundleCart) Update(tx *gorm.DB, data Cart) (results Cart, err error) {
	err = tx.Model(&results).
		Updates(map[string]interface{}{
			"sku":       data.Sku,
			"client_id": data.ClientId,
			"quantity":  data.Quantity,
			"is_active": data.IsActive,
		}).
		Where("client_id = ?", data.ClientId).
		Where("sku = ?", data.Sku).
		Error

	if err != nil {
		tx.Rollback()
		return results, err
	}

	return results, err
}
