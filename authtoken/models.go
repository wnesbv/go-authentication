package authtoken

import (
    "github.com/golang-jwt/jwt/v5"
)

type Claims struct {
    User_id int `json:"user_id"`
    Email string `json:"email"`
    jwt.RegisteredClaims
}
