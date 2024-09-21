package xjwt_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/avila-r/xgo/pkg/xjwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Test_Generate(t *testing.T) {
	x := &xjwt.S{
		Secret: "my-secret",
	}

	token, err := x.Generate("1234")

	if err != nil {
		t.Fatal(err)
	}

	if token == "" {
		t.Fatalf("Generated token is empty")
	}

	parsed, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(x.Secret), nil
	})

	if err != nil {
		t.Fatalf("Error parsing token: %v", err)
	}

	if !parsed.Valid {
		t.Fatal("Generated token is invalid")
	}

	claims, ok := parsed.Claims.(jwt.MapClaims)

	if !ok {
		t.Fatal("Failed to extract claims from token")
	}

	if claims["id"] != "1234" {
		t.Errorf("Expected ID '1234', but got '%v'", claims["id"])
	}

	if claims["admin"] != true {
		t.Errorf("Expected 'admin' to be true, but got '%v'", claims["admin"])
	}

	// Check if the expiration time is correct
	if exp, ok := claims["exp"].(float64); !ok || exp < float64(time.Now().Unix()) {
		t.Errorf("Token expired or 'exp' claim is invalid")
	}
}

func Test_ValidToken(t *testing.T) {
	app := fiber.New()

	x := &xjwt.S{"my-secret-key"}

	app.Use(x.DefaultMiddleware())

	app.Get("/protected", func(c *fiber.Ctx) error {
		return c.SendString("allowed access")
	})

	token, _ := x.Generate("user-id")

	request := httptest.NewRequest(http.MethodGet, "/protected", nil)

	request.Header.Set("Authorization", "Bearer "+token)

	response, _ := app.Test(request)

	code := response.StatusCode

	// Expects Status OK
	if code != 200 {
		t.Fatalf("Expected status %d, received %d", 200, code)
	}
}

func Test_InvalidToken(t *testing.T) {
	app := fiber.New()

	x := &xjwt.S{"my-secret-key"}

	app.Use(x.DefaultMiddleware())

	app.Get("/protected", func(c *fiber.Ctx) error {
		return c.SendString("allowed access")
	})

	token := "invalid token"

	request := httptest.NewRequest(http.MethodGet, "/protected", nil)

	request.Header.Set("Authorization", "Bearer "+token)

	response, _ := app.Test(request)

	code := response.StatusCode

	// Expects Status Unauthorized
	if code != 401 {
		t.Fatalf("Expected status %d, received %d", 401, code)
	}
}
