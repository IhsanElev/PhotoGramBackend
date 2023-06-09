package helpers

import (
	"log"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

var secretKey = os.Getenv("SECRET_KEY")

func GenerateToken(id uint, email string, role string) string {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
		"role":  role,
	}
	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := parseToken.SignedString([]byte(secretKey))
	return signedToken
}
