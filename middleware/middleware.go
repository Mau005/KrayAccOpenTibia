package middleware

import (
	"net/http"
	"strings"

	"github.com/Mau005/KrayAccOpenTibia/models"
	"github.com/Mau005/KrayAccOpenTibia/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, utils.ErrorAuthorizationMission, http.StatusUnauthorized)
			return
		}

		tokenParts := strings.Split(tokenString, "")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			http.Error(w, utils.ErrorInvalidTokenFormat, http.StatusUnauthorized)
		}
		claims := &models.Claim{}
		token, err := jwt.ParseWithClaims(tokenParts[1], claims, func(t *jwt.Token) (interface{}, error) {
			return utils.PasswordSecurityDefaul, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, utils.ErrorInvalidToken, http.StatusUnauthorized)
			return
		}
		context.Set(r, utils.CtxAccountName, claims.AccountName)
		context.Set(r, utils.CtxAccountEmail, claims.Email)
		context.Set(r, utils.CtxAccountID, claims.IDAccount)
		next.ServeHTTP(w, r)
	})
}
