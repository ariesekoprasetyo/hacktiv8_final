package controller

import (
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"time"
)

func GenerateToken(userId uint) (string, error) {
	var jwtSecret GenerateTokenStruct
	token := jwt.New(jwt.SigningMethodHS256)
	now := time.Now().UTC()
	claims := token.Claims.(jwt.MapClaims)

	claims["sub"] = userId
	claims["exp"] = now.Add(time.Hour * 24 * 2).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	tokenString, err := token.SignedString([]byte(jwtSecret.JWTSecret))
	if err != nil {
		return "", fmt.Errorf("generetaing JWT Token Failed: %w", err)
	}
	return tokenString, nil
}

func ValidateToken(token string, signedJWTKey string) (interface{}, error) {
	tok, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
		}

		return []byte(signedJWTKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("invalidate token: %w", err)
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		return nil, fmt.Errorf("invalid token claim")
	}

	return claims["sub"], nil
}
