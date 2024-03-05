package authenticator

import (
	"fmt"
	"gintut/helpers/controllers"
	"gintut/initializers"
	"gintut/models"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type MyCustomClaims struct {
	jwt.RegisteredClaims
}

// get personal data by their NIK
func GetPersonalByNIK(nik string) (models.TPersonal, error) {
	var personal models.TPersonal
	var field = fmt.Sprintf(`"NIK"`)
	err := controllers.GetModel(&personal, field, nik)
	return personal, err
}

// generate jwtn
//
// sub: nik
//
// nam: name
//
// eml: email
func GenerateJWT(nik string) (string, error) {
	initializers.InitEnv()
	lifetime := os.Getenv("TOKEN_LIFE_HOUR")
	secret := os.Getenv("SUPER_SECRET")

	life, err := strconv.Atoi(lifetime)
	if err != nil {
		fmt.Println("Error converting value to int:", err)
		return "", err
	}

	personal, err := GetPersonalByNIK(nik)
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
