package models

import "github.com/dgrijalva/jwt-go"

type Claim struct {
	IDAccount   int    `json:"id"`
	AccountName string `json:"username"`
	Email       string `json:"Email"`
	TypeAccess  uint   `json:"typeaccess"`
	jwt.StandardClaims
}
