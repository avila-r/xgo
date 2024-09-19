package xjwt

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var (
	key *rsa.PrivateKey
)

var (
	default_config = jwtware.Config{
		KeyFunc: func(t *jwt.Token) (interface{}, error) {
			if t.Method.Alg() != jwtware.HS256 {
				return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
			}

			if key == nil {
				key, _ = rsa.GenerateKey(rand.Reader, 2048)
			}

			return key.Public(), nil
		},
	}
)

var (
	DefaultMiddleware = jwtware.New(default_config)

	Middleware = func(config ...jwtware.Config) func(*fiber.Ctx) error {
		return jwtware.New(config...)
	}
)
