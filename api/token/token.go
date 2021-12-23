package token

import (
	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	Username  string
	Authority int
	jwt.StandardClaims
}

func GenerateToken(username string) string {
	tk := &Token{Username: username, Authority: 0}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte("secert"))

	return tokenString
}

func ValidToken(token string) (*Token, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Token{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secert"), nil
	})

	if err == nil && tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Token); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
