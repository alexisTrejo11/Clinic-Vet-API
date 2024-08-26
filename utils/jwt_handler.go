package utils

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var JWTSecret = []byte("your_secret_key")

type Claims struct {
	UserID string   `json:"user_id"`
	Roles  []string `json:"roles"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID string, roles []string) (string, error) {
	expirationTime := time.Now().Add(3 * time.Hour)

	claims := &Claims{
		UserID: userID,
		Roles:  roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Create the token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	return token.SignedString(JWTSecret)
}

func ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return JWTSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func ExtractTokenFromHeader(c *fiber.Ctx) (string, error) {
	// Obtener el valor del encabezado Authorization
	authHeader := c.Get("Authorization")

	// Verificar si el encabezado Authorization está presente
	if authHeader == "" {
		return "", errors.New("missing Authorization header")
	}

	// Eliminar el prefijo "Bearer " para obtener solo el token
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("invalid Authorization header format")
	}

	return parts[1], nil
}

func GetUserIDFromRequest(c *fiber.Ctx) (int, error) {
	tokenString, err := ExtractTokenFromHeader(c)
	if err != nil {
		return 0, err
	}

	claims, err := ValidateJWT(tokenString)
	if err != nil {
		return 0, err
	}

	userID, err := strconv.Atoi(claims.UserID)
	if err != nil {
		return 0, fiber.NewError(fiber.StatusBadRequest, "Can't Process Id")
	}
	return userID, nil
}
