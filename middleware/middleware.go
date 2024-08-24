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

		var tokenString string

		if cookie, err := r.Cookie(utils.NameCookieToken); err == nil {
			tokenString = cookie.Value
		} else {
			tokenString = r.Header.Get("Authorization")
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		if tokenString == "" {
			http.Error(w, "Missing or invalid token", http.StatusUnauthorized)
			return
		}

		claims := &models.Claim{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(utils.PasswordSecurityDefaul), nil
		})
		if err != nil || !token.Valid {
			http.Error(w, utils.ErrorInvalidToken, http.StatusUnauthorized)
			return
		}
		context.Set(r, utils.CtxAccountName, claims.AccountName)
		context.Set(r, utils.CtxAccountEmail, claims.Email)
		context.Set(r, utils.CtxAccountID, claims.AccountID)
		context.Set(r, utils.CtxTypeAccount, claims.TypeAccess)
		context.Set(r, utils.CtxClaim, *claims)
		next.ServeHTTP(w, r)
	})
}

func CommonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Expose-Headers", "Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, thorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		next.ServeHTTP(w, r)
	})
}
