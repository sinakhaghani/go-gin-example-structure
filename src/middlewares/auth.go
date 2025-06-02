package middlewares

import (
    "log"
    "net/http"
    "os"
    "strings"

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
    "github.com/joho/godotenv"
)

var JwtKey []byte

func AuthMiddleware() gin.HandlerFunc {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("❌ .env file not loaded:", err)
    }

    secret := os.Getenv("SECRET_KEY")
    if secret == "" {
        log.Fatal("❌ SECRET_KEY is missing in .env")
    }

    JwtKey = []byte(secret)

    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
            c.Abort()
            return
        }

        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
            c.Abort()
            return
        }

        tokenStr := parts[1]
        token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
            return JwtKey, nil
        })

        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
            c.Abort()
            return
        }

        c.Next()
    }
}
