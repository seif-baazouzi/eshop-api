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

## Shop Routes

- GET     /shops
- GET     /shops/user
- POST    /shops            { shopName, shopDescription }
- PUT     /shops/:shopName  { shopName, shopDescription }
- PATCH   /shops/:shopName  { image }
- DELETE  /shops/:shopName
- PUT     /shops/:shopName/rate { rate }

## Items Routes

- GET     /items
- GET     /items/shop/:shopName
- GET     /items/:itemID
- POST    /items/:shopName  { itemName, itemDescription, itemPrice }
- PUT     /items/:itemID    { itemName, itemDescription, itemPrice }
- PATCH   /items/:itemID    { image }
- DELETE  /items/:itemID
- PUT     /items/:itemID/rate { rate }

## Comments Routes

- GET     /comments/:itemID
- POST    /comments/:itemID       { commentValue }
