// server/jwt/jwt.go

package jwt

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

var errENV error

var signingKey = []byte(os.Getenv("JWT_SECRET_TOKEN"))
var tokenCookieName = "auth_token"

type Claims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

func init() {
	errENV = godotenv.Load()
	if errENV != nil {
		log.Fatal("Error loading .env file: ", errENV)
	}
}

func GenerateJWTToken(userID int, expiration time.Duration) (string, error) {
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expiration).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}

	return tokenString, err
}

func RemoveJWTTokenFromCookie() *http.Cookie {
	return &http.Cookie{
		Name:     tokenCookieName,
		Value:    "",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
	}
}

func SetCookieWithJWTToken(w http.ResponseWriter, token string, expiration time.Duration) {
	cookie := &http.Cookie{
		Name:     tokenCookieName,
		Value:    token,
		Expires:  time.Now().Add(expiration),
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	fmt.Println("Cookie Name:", cookie.Name)
	fmt.Println("Cookie Value:", cookie.Value)
}

func VerifyJWTToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}
