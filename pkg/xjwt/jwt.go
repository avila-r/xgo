package xjwt

import (
	"fmt"
	"time"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type S struct {
	Secret string
}

func (x *S) DefaultMiddleware() func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		KeyFunc: func(t *jwt.Token) (interface{}, error) {
			if t.Method.Alg() != jwtware.HS256 {
				return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
			}

			return []byte(x.Secret), nil
		},
	})
}

func (x *S) Middleware(config ...jwtware.Config) func(*fiber.Ctx) error {
	return jwtware.New(config...)
}

func (x *S) Generate(id string) (string, error) {
	claims := jwt.MapClaims{
		"id":    id,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
		"admin": true,
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := t.SignedString([]byte(x.Secret))

	if err != nil {
		return "", err
	}

	return token, nil
}
