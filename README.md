# Check Out

Simple check out process with promo.

## Feature

Query
```bash
product(sku: String!): Product!
allProducts(param: PagingQuery!): Products!
promo(sku: String!): Promo!
allActivePromo(param: PagingQuery!): Promos!
cartList(param: PagingQuery!): Carts!
checkout(input: CheckOutItem): ResponseCheckout!
```

Mutation
```bash
createProduct(input: NewProduct): Product!
createPromo(input: NewPromo): Promo!
updateProduct(input: NewProduct): Product!
updatePromo(input: NewPromo): Promo!
addCart(input: AddToCartItem!): Cart!
```

## Requirement
- direnv _(optional)_
- `make >= v4.1`  you need this tool to use the `Makefile`
- `fswatch >= v1.11.2` if you want to use auto-compile-restart feature on `Makefile`
- golang 1.16.x
- docker / podman
- PostgreSQL >= 10

## Development

1. First time only:

```shell
cp .envrc.sample .envrc
````

2. Do your adjustment on `.envrc`, then

```shell
direnv allow .
```

do this each time you change value inside the `.envrc`

3. run server on auto-reload mode

```shell
make serve
```

## Documentation

After running, 
- GraphQL playground will `http://127.0.0.1:2727/v1/graphql`
- Endpoint will serve in `http://127.0.0.1:2727/v1/graphql/query`

All documentation for Query and Mutation will serve in GraphQL playground.

## Available Makefile Command

```bash
[!] Available Command: 
-----------------------
app-image                      Create a docker image
compile                        Build binary version of application
compile_dev                    Compile dev version of application
lint                           Lint this codebase
run                            execute go run main.go
serve                          Run application and automaticaly restart on source code change
```