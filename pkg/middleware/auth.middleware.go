package middleware

import (
	"Codimite_Assignment/internal/models"
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type contextKey string

const (
	UserContextKey     contextKey = "userID"
	UserNameContextKey contextKey = "userName"
)

// AuthMiddleware for authenticate requests
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
		// Load environment variables
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		// get the JWT secret key from env
		secret := os.Getenv("JWT_SECRET_KEY")

		// get the Authorization header from the request
		authHeader := r.Header.Get("Authorization")

		// Check if Authorization header format
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Missing Token", http.StatusUnauthorized)
			return
		}

		// Extract token string by removing the "Bearer " prefix
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims := &models.UserToken{}
		token, errToken := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		log.Println("Claims", tokenString)
		
		if errToken != nil || !token.Valid {
			log.Println("Error token",errToken)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Store the extracted data from the token in context
		ctx := context.WithValue(r.Context(), UserContextKey, claims.UserId)
		ctx = context.WithValue(ctx, UserNameContextKey, claims.UserName)
		
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
