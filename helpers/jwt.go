package helpers

import (
	"fmt"
	"log"
	"os"

	m "github.com/LordFeexz/Nusa-Tech/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

func getSecret() (string, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error")
		return "", err
	}

	secret := os.Getenv("SECRET")

	return secret, nil
}

func CreateToken(user m.User) (string, error) {

	data := map[string]interface{}{
		"id":    user.Id,
		"email": user.Email,
	}

	secret, _ := getSecret()

	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{"data": data})

	tokenStr, err := token.SignedString(secret)

	if err != nil {
		fmt.Println("error", err.Error())
		return "", err
	}

	return tokenStr, nil
}

func Validate(tokenStr string) bool {
	secret, _ := getSecret()

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return false
	}

	if token.Valid {
		return true
	} else {
		return false
	}
}
