package midleware

import (
	"MiniProject/constant"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(userID int, name string, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	claims["name"] = name
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constant.SECRET_JWT))
}
