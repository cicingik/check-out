package database

import (
	"context"
	"fmt"
	"gorm.io/gorm/clause"

	"gorm.io/gorm"
)

type (
	Product struct {
		Model
		Sku      string  `gorm:"Column:sku" json:"sku"`
		Name     string  `gorm:"Column:name" json:"name"`
		Price    float64 `gorm:"Column:price" json:"price"`
		Quantity int     `gorm:"Column:quantity" json:"quantity"`
	}

	ProductPagination struct {
		Product []Product `json:"product"`
		Pg      struct {
			After     int `json:"after"`
			Before    int `json:"before"`
			TotalPage int `json:"total"`
			Page      int `json:"page"`
		} `json:"pagination"`
	}

	BundleProduct struct {
		db *gorm.DB
		t  Product
	}
)

func (t *Product) TableName() string {
	return "ecommerce.products"
}

func InitProduct(ctx context.Context, g *gorm.DB) *BundleProduct {
	return &BundleProduct{
		db: g.WithContext(ctx),
		t:  Product{},
	}
}

func (b *BundleProduct) Create(data Product) (Product, error) {
	var err error

	err = b.db.Create(&data).Error
	if err != nil {
		return Product{}, err
	}

	return data, err
}

func (b *BundleProduct) FindAll(data RequestPaginatorParam) (r ProductPagination, err error) {

	var product []Product

	tx := b.db.
		Model(Promo{}).
		Select(`"ecommerce"."product".*`)

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
	}, &product)

	r.Product = product
	r.Pg.After = paging.NextPage
	r.Pg.Before = paging.PrevPage
	r.Pg.TotalPage = paging.TotalPage
	r.Pg.Page = data.Page

	return r, nil
}

func (b *BundleProduct) FindProduct(column string, value string) (Product, error) {
	var err error
	var product Product

	query := fmt.Sprintf(`%s = ?`, column)
	err = b.db.Model(Product{}).Where(query, value).First(&product).Error
	if err != nil {
		return product, err
	}

	return product, err
}

func (b *BundleProduct) FindProductStock(skus []string) (results []Product, err error) {

	// use locking for other purchase process
	err = b.db.Clauses(clause.Locking{
		Strength: "UPDATE",
	}).Find(&results).
		Where(`sku in ?`, skus).
		Error

	if err != nil {
		return results, err
	}

	return results, nil
}

func (b *BundleProduct) Update(tx *gorm.DB, data Product) (results Product, err error) {
	err = tx.Model(&results).
		First(&results, data.ID).
		Updates(map[string]interface{}{
			"quantity": data.Quantity,
			"sku":      data.Sku,
			"name":     data.Name,
			"price":    data.Price,
		}).
		Error

	if err != nil {
		tx.Rollback()
		return results, err
	}

	return results, err
}
