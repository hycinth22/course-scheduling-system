package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"log"
	"time"
)

type AuthInfo struct {
	UserID int `json:"uid"`
	jwt.RegisteredClaims
}

var (
	hmacSecret = "webcourseschedulinghmacSampleSecret"
)

func SignToken(uid int) (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, AuthInfo{
		UserID: uid,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "coursescheduling",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                         // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                         // 生效时间
		},
	})
	return token.SignedString(hmacSecret)
}

func ParseToken(tokenString string) (auth *AuthInfo, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &AuthInfo{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})
	if claims, ok := token.Claims.(*AuthInfo); ok && token.Valid {
		log.Printf("%v %v", claims.UserID, claims.RegisteredClaims.IssuedAt)
		return claims, nil
	} else {
		log.Println(err)
		return nil, err
	}
}
