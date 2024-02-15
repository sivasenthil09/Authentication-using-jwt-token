package models

import "github.com/dgrijalva/jwt-go"

type Credentials struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
