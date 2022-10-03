package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"

	"golang-mysql-restful-starter-kit/config"
	"golang-mysql-restful-starter-kit/helpers"
)

func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] == nil {
			helpers.SendRespond(w, http.StatusUnauthorized, nil, helpers.ErrUnauthorized)
			return
		}

		var mySigningKey = []byte(os.Getenv(config.JWT_SECRET))

		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf(helpers.ErrUnauthorized)
			}
			return mySigningKey, nil
		})

		if err != nil {
			helpers.SendRespond(w, http.StatusUnauthorized, nil, helpers.ErrUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			r.Header.Set("LoggedInUser", fmt.Sprintf("%v", claims["userId"]))
			handler.ServeHTTP(w, r)
			return
		}
		helpers.SendRespond(w, http.StatusUnauthorized, nil, helpers.ErrUnauthorized)
	}
}
