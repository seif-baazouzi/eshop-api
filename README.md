# E-shop REST API

This is a rest api for e-shop stores, were the user can create his own shops and managed.

# Used Technologies

- GoLang
- GoFiber
- Postgresql
- Redis

# Quick Start

```console
$ export DB_USER=postgres
$ export DB_PASSWORD=password
$ export DB_HOST=127.0.0.1
$ export DB_NAME=eshop
$ export REDIS_HOSTNAME=127.0.0.1
$ export JWT_SECRET=JWT_SECRET
$ export UPLOADING_DIRECTORY='/path/to/uploading/directory'

$ go run ./src/main.go
```

# Routes

## Auth routes

- POST    /login    { email, password }
- POST    /signup   { email, username, password }
- GET     /username

## Shop Routes

- GET     /shops?page=pageNumber
- GET     /shops/user?page=pageNumber
- GET     /shops/:shopName
- GET     /shops/user/rate/:shopName
- POST    /shops            { shopName, shopDescription }
- PUT     /shops/:shopName  { shopName, shopDescription }
- PATCH   /shops/:shopName  { image }
- DELETE  /shops/:shopName
- PUT     /shops/:shopName/rate { rate }

## Items Routes

- GET     /items?page=pageNumber
- GET     /items/shop/:shopName?page=pageNumber
- GET     /items/:itemID
- GET     /items/user/rate/:itemID
- POST    /items/:shopName  { itemName, itemDescription, itemPrice }
- PUT     /items/:itemID    { itemName, itemDescription, itemPrice }
- PATCH   /items/:itemID    { image }
- DELETE  /items/:itemID
- PUT     /items/:itemID/rate { rate }

## Comments Routes

- GET     /comments/:itemID?page=pageNumber
- POST    /comments/:itemID       { commentValue }
- POST    /comments/:commentID    { commentValue }
- DELETE  /comments/:commentID

## Carts Routes

- GET     /carts/shop/:shopName?page=pageNumber
- GET     /carts/user?page=pageNumber
- GET     /carts/shop/items/:cartID
- GET     /carts/user/items/:cartID
- POST    /carts  { address, shopName, items }
- PUT     /carts/:cartID/:status
