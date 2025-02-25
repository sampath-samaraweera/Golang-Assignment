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

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		secret := os.Getenv("JWT_SECRET_KEY")
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Missing Token", http.StatusUnauthorized)
			return
		}
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
