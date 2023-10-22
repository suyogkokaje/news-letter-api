package auth

import (
    "github.com/dgrijalva/jwt-go"
    "time"
)

var (
    jwtSecret = []byte("your-secret-key")
)

// Claims represents the JWT claims structure.
type Claims struct {
    UserID   uint   `json:"user_id"`
    Role     string `json:"role"`
    jwt.StandardClaims
}

func GenerateToken(userID uint, role string) (string, error) {
    claims := &Claims{
        UserID:   userID,
        Role:     role,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().Add(7 * 24 * time.Hour).Unix(),
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}
