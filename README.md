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
- PUT     /shops/:shopName/rate { rate }
