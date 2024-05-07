package auth

import (
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

var (
	SecretKey = "your_secret_key"
	RoleAdmin = "admin"
	RoleUser  = "user"
)

type MyClaims struct {
	jwt.StandardClaims
	UserID string `json:"user_id"`
}

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "No token provided")
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})
		if err != nil || !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or expired token")
		}
		return next(c)
	}
}

func GenerateToken(userID uint, role string) (tokenString string, err error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour).Unix()

	tokenString, err = token.SignedString([]byte(SecretKey))
	if err != nil {
		return tokenString, err
	}
	return tokenString, err
}

func GetTokenClaims(c echo.Context) (claims jwt.MapClaims, err error) {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return claims, err
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	claims, _ = token.Claims.(jwt.MapClaims)
	return claims, err
}
