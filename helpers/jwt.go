package helpers

import (
	"golang-mysql-restful-starter-kit/config"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJWT(ID int32) (string, error) {

	jwtToken := jwt.New(jwt.SigningMethodHS256)

	claims := jwtToken.Claims.(jwt.MapClaims)
	claims["userId"] = ID
	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()

	token, err := jwtToken.SignedString([]byte(os.Getenv(config.JWT_SECRET)))

	if err != nil {
		log.Printf("JWT token error -  %s", err.Error())
		return "", err
	}
	return token, nil
}
