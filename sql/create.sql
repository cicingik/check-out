CREATE SCHEMA IF NOT EXISTS ecommerce;

CREATE TABLE IF NOT EXISTS ecommerce.carts
(
    id         bigserial primary key,
    sku        varchar(50) not null,
    quantity   int         not null,
    status     int         not null,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    deleted_at timestamp
    );

CREATE TABLE IF NOT EXISTS ecommerce.products
(
    id         bigserial primary key,
    sku        varchar(100) not null unique,
    name       varchar(100) not null,
    price      float        not null,
    quantity   int          not null,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    deleted_at timestamp
    );

CREATE TABLE IF NOT EXISTS ecommerce.promo
(
    id                bigserial primary key,
    sku               varchar(100) not null unique,
    promo_type        varchar(100) not null,
    minimal_purchased int,
    bonus_product_sku varchar(100),
    discount          float,
    created_at        timestamp default now(),
    updated_at        timestamp default now(),
    deleted_at        timestamp
    );

DROP TYPE IF EXISTS ecommerce.promo_type;
CREATE TYPE ecommerce.promo_type AS ENUM ('discount', 'free_item');

INSERT INTO ecommerce.products (sku, name, price, quantity)
VALUES ('120P90', 'Google Home', 49.99, 10),
       ('43N23P', 'MacBook Pro', 5399.99, 5),
       ('A304SD', 'Alexa Speaker', 109.50, 10),
       ('234123', 'Raspberry Pi B', 30.00, 2)
    ON CONFLICT (sku) DO UPDATE
                             SET name     = EXCLUDED.name,
                             price    = EXCLUDED.price,
                             quantity = EXCLUDED.quantity;

INSERT INTO ecommerce.promo (sku, promo_type, minimal_purchased, bonus_product_sku, discount)
VALUES ('43N23P', 'free_item', 1, '234123', 0.0),
       ('120P90', 'free_item', 2, '120P90', 0.0),
       ('A304SD', 'discount', 3, null, 10.0)
    ON CONFLICT (sku) DO UPDATE
                             SET promo_type        = EXCLUDED.promo_type,
                             minimal_purchased = EXCLUDED.minimal_purchased,
                             bonus_product_sku = EXCLUDED.bonus_product_sku,
                             discount          = EXCLUDED.discount;

