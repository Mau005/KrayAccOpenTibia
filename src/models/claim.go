package models

import "github.com/dgrijalva/jwt-go"

type Claim struct {
	AccountID   int    `json:"id"`
	AccountName string `json:"username"`
	Email       string `json:"Email"`
	TypeAccess  uint8  `json:"typeaccess"`
	jwt.StandardClaims
}
