package api

import (
	"github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
)

var (
	RecommendedConfig = fiber.Config{
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
		ErrorHandler: ErrorHandler,
	}
)
