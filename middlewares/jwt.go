package middlewares

import (
    "net/http"
    "os"
    "strings"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
    "github.com/joho/godotenv"
)

var jwtKey []byte

func init() {
    // Load .env saat package di-load
    _ = godotenv.Load()
    jwtKey = []byte(os.Getenv("JWT_SECRET"))
}

type JWTClaim struct {
    UserID uint   `json:"user_id"`
    Role   string `json:"role"`
    jwt.RegisteredClaims
}

// Generate token saat login
func GenerateJWT(userID uint, role string) (string, error) {
    claims := &JWTClaim{
        UserID: userID,
        Role:   role,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}

// Middleware autentikasi token
func JWTAuth() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenStr := getTokenFromHeader(c)
        if tokenStr == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak ditemukan"})
            c.Abort()
            return
        }

        claims := &JWTClaim{}
        token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
            return jwtKey, nil
        })

        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
            c.Abort()
            return
        }

        // Simpan data user di context
        c.Set("userID", claims.UserID)
        c.Set("role", claims.Role)
        c.Next()
    }
}

// Middleware khusus admin
func AdminOnly() gin.HandlerFunc {
    return func(c *gin.Context) {
        role, exists := c.Get("role")
        if !exists || role != "admin" {
            c.JSON(http.StatusForbidden, gin.H{"error": "Hanya admin yang diizinkan"})
            c.Abort()
            return
        }
        c.Next()
    }
}

// Ambil token dari header Authorization
func getTokenFromHeader(c *gin.Context) string {
    authHeader := c.GetHeader("Authorization")
    if strings.HasPrefix(authHeader, "Bearer ") {
        return strings.TrimPrefix(authHeader, "Bearer ")
    }
    return ""
}
