package auth

import (
	"Codimite_Assignment/internal/models"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func GenerateToken(user_id int, username string) (string, error) {
	// Load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	expirationTime := time.Now().Add(24*time.Hour)

	// Define token claims with user info
	claims := models.UserToken{
		UserId: user_id,
		UserName: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Get the JWT secret key from env
	secret := os.Getenv("JWT_SECRET_KEY")
	var jwtSecret = []byte(secret)

	// Create a new JWT token 
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	tokenStr, err := token.SignedString(jwtSecret)
	if err != nil {
		log.Println("Error signing token: ",token)
		return "", err
	}

	return tokenStr, nil
}