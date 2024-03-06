package authenticator

import (
	"fmt"
	"gintut/helpers/controllers"
	"gintut/initializers"
	"gintut/models"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type MyCustomClaims struct {
	jwt.RegisteredClaims
}

// get personal data by their NIK
func GetPersonalDetail(field string, value string) (models.TPersonal, error) {
	var personal models.TPersonal
	var lookup = fmt.Sprintf(`"%s"`, strings.ToUpper(field))
	err := controllers.GetModel(&personal, lookup, value)
	return personal, err
}

// generate jwt
func GenerateJWT(nik string) (string, error) {
	initializers.InitEnv()
	lifetime := os.Getenv("TOKEN_LIFE_HOUR")
	secret := os.Getenv("SUPER_SECRET")

	life, err := strconv.Atoi(lifetime)
	if err != nil {
		fmt.Println("Error converting value to int:", err)
		return "", err
	}

	personal, err := GetPersonalDetail("nik", nik)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	claims := jwt.MapClaims{
		"exp":  jwt.NewNumericDate(time.Now().Add(time.Duration(life) * time.Hour)),
		"iss":  "system",
		"sub":  personal.Nik,
		"eml":  personal.Eml,
		"nam":  personal.Nam,
		"rlcd": personal.Rlcd,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return ss, nil
}
