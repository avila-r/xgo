package api

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

type Error struct {
	Code    int
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

func NewError(code int, message string) *Error {
	return &Error{Code: code, Message: message}
}

var (
	ErrorHandler = func(c *fiber.Ctx, err error) error {
		c.Set(
			fiber.HeaderContentType,
			fiber.MIMETextPlainCharsetUTF8,
		)

		// Tries to parse 'error' as 'APIError'
		var e *Error

		if !errors.As(err, &e) {
			// Status: Internal Server Error [500]
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		var (
			code     = e.Code
			code_str = utils.StatusMessage(code)
		)

		// If an invalid status code is provided,
		// replaces with [500] Internal Server Error.
		if e.Code == 0 || utils.StatusMessage(code) == "" {
			code = fiber.StatusInternalServerError
			code_str = "Internal Server Error"
		}

		return c.Status(code).JSON(fiber.Map{
			"code":    code_str,
			"message": e.Message,
		})
	}
)
