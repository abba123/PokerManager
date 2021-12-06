package token

import (
	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	Username string
	jwt.StandardClaims
}

func GenerateToken(username string) string {
	tk := &Token{Username: username}
	tk.StandardClaims.Subject = username
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte("secert"))

	return tokenString
}

func ValidToken(token string) (*Token, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Token{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secert"), nil
	})
	if err != nil {
		return nil, err
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Token); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
