package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/rulanugrh/uranus/configs"
	"github.com/rulanugrh/uranus/internal/domain/entity"
	"github.com/rulanugrh/uranus/internal/domain/web"
)

type jwtClaims struct {
	Email string
	Name  string
	jwt.RegisteredClaims
}

func GenerateToken(user entity.User) (string, error) {
	conf := configs.GetConfig()
	time := jwt.NewNumericDate(time.Now().Add(1 * time.Hour))

	claims := &jwtClaims{
		Email: user.Email,
		Name:  user.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: time,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(conf.JWTSecret))
	if err != nil {
		log.Printf("Cant claim jwt token: %v", err)
	}

	return tokenString, nil
}

func ValidateToken(token string) error {
	conf := configs.GetConfig()
	tokens, err := jwt.ParseWithClaims(token, &jwtClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(conf.JWTSecret), nil
	})

	if err != nil {
		log.Printf("Token is not valid: %v", err)
	}

	claims, errClaim := tokens.Claims.(*jwtClaims)
	if !errClaim {
		log.Printf("Cant claim token %v", errClaim)
	}

	if claims.ExpiresAt.Unix() < time.Now().Local().Unix() {
		log.Printf("token expired")
	}

	return nil

}

func JWTVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var token = r.Header.Get("Authorization")
		json.NewEncoder(w).Encode(r)

		token = strings.TrimSpace(token)
		if token == "" {
			res := web.ResponseFailure{
				Code:    http.StatusUnauthorized,
				Message: "cant login because not have token",
			}

			response, _ := json.Marshal(res)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(response)
		}

		err := ValidateToken(token)
		if err != nil {
			res := web.ResponseFailure{
				Code:    http.StatusUnauthorized,
				Message: "Cant login because token is not valid",
			}

			response, _ := json.Marshal(res)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(response)
		}

		next.ServeHTTP(w, r)

	})
}
