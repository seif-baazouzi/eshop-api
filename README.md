# E-shop REST API

This is a rest api for e-shop stores, were the user can create his own shops and managed.

# Used Technologies

- GoLang
- GoFiber
- Postgresql
- Redis

# Routes

## Auth routes

- POST    /login    { email, password }
- POST    /signup   { email, username, password }
- GET     /username

## Shop Routes

- GET     /shops?page=pageNumber
- GET     /shops/user?page=pageNumber
- GET     /shops/:shopName
- POST    /shops            { shopName, shopDescription }
- PUT     /shops/:shopName  { shopName, shopDescription }
- PATCH   /shops/:shopName  { image }
- DELETE  /shops/:shopName
- PUT     /shops/:shopName/rate { rate }

## Items Routes

- GET     /items?page=pageNumber
- GET     /items/shop/:shopName?page=pageNumber
- GET     /items/:itemID
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
