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
- POST    /shops            { name, description }
- PUT     /shops/:shopName  { name, description }
- PATCH   /shops/:shopName  { image }
- DELETE  /shops/:shopName
- PUT     /shops/:shopName/rate { rate }

## Items Routes

- GET     /items
- GET     /items/:shopName
