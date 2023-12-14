package auth

import (
	"errors"
	"time"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var SECRET_KEY string

type Claims struct {
	Id primitive.ObjectID
	jwt.StandardClaims
}

func EncodeJWT(id primitive.ObjectID, minutes time.Duration) (*string, error) {
	var claims *Claims = &Claims{
		Id: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(minutes * time.Minute).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "skycrypt-api",
		},
	}

	var token *jwt.Token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(SECRET_KEY))

	if err != nil {
		return nil, err
	}

	return &tokenStr, nil
}

func ValidateJWT(tokenString string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(SECRET_KEY), nil 
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return &claims, nil
	}
	
	return nil, errors.New("invalid token")
}
