package JWT

import (
	"go-chatbot/Models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	TokenAExpireDuration = time.Minute * 10
	TokenRExpireDuration = time.Hour * 24 * 7
)

type MyClaims struct {
	UserID int64 `json:"user_id"`
	jwt.StandardClaims
}

func GenToken(userID int64) (aToken, rToken string, err error) {
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, MyClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenAExpireDuration).Unix(),
			Issuer:    "my-project",
		},
	}).SignedString(Models.JWTSercet)
	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(TokenRExpireDuration).Unix(),
		Issuer:    "my-project",
	}).SignedString(Models.JWTSercet)
	return aToken, rToken, err
}

func ParseToken(tokenString string) (*MyClaims, error) {
	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, JWTKeyFunc)
	if err != nil {
		return nil, err
	}
	if token.Valid {
		return mc, nil
	}
	return nil, Models.ErrorAuthHeaderValid
}

func ParseRToken(rToken string) (err error) {
	var mc = new(jwt.StandardClaims)
	token, err := jwt.ParseWithClaims(rToken, mc, JWTKeyFunc)
	if err != nil {
		return err
	}
	if token.Valid {
		return nil
	}
	return Models.ErrorAuthHeaderValid
}

func RefreshToken(aToken, rToken string) (newAToken, newRtoken string, err error) {
	if err = ParseRToken(rToken); err != nil {
		return
	}
	var mc = new(MyClaims)
	_, err = jwt.ParseWithClaims(aToken, mc, JWTKeyFunc)
	v, _ := err.(jwt.ValidationError)
	if v.Errors == jwt.ValidationErrorExpired {
		return GenToken(mc.UserID)
	}
	return
}

func JWTKeyFunc(token *jwt.Token) (interface{}, error) {
	return Models.JWTSercet, nil
}
