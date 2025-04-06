package auth

import (
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

var jwtTTL = 24 * time.Hour
var jwtSecret []byte

func init() {
	err := godotenv.Load()

	if err != nil {
		panic("Error loading JWT secret")
	}

	jwtSecret = []byte(os.Getenv("JWT_SECRET"))
}

func GenerateJWT(id uuid.UUID) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject:   id.String(), // User ID (as string)
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(jwtTTL)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	})
	jwtString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return jwtString, nil
}

func ValidateJWT(jwtToken string) (uuid.UUID, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		return uuid.Nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return uuid.Nil, err
	}

	if _, ok := claims["sub"]; !ok {
		return uuid.Nil, err
	}

	id, err := uuid.Parse(claims["sub"].(string))
	if err != nil {
		return uuid.Nil, err
	}
	return id, err
}

func SetAuthCookie(w http.ResponseWriter, jwtToken string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    jwtToken,
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Now().Add(24 * time.Hour),
	})
}

func ClearAuthCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    "",
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
	})
}
