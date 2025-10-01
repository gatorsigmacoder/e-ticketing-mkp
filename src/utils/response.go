package utils

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Status   string      `json:"status"`  // e.g., "success" or "error"
	Message  string      `json:"message"` // A short description
	Result   interface{} `json:"result"`  // The payload (can be nil for errors)
	MetaData interface{} `json:"meta_data"`
}

func NewResponse(status, message string, data interface{}, metaData interface{}) *Response {
	return &Response{
		Status:   status,
		Message:  message,
		Result:   data,
		MetaData: metaData,
	}
}

func FormatValidationError(err error) map[string]string {
	errors := make(map[string]string)
	for _, err := range err.(validator.ValidationErrors) {
		errors[err.Field()] = err.Tag()
	}
	return errors
}

func Success(c *fiber.Ctx, statusCode int, message string, data interface{}, metaData interface{}) error {
	return c.Status(statusCode).JSON(NewResponse("success", message, data, metaData))
}

func Error(c *fiber.Ctx, statusCode int, message string, data interface{}, metaData interface{}) error {
	return c.Status(statusCode).JSON(NewResponse("error", message, data, metaData))
}

func PaginationResponse(data interface{}, total int64, page int, per_page int) map[string]interface{} {
	result := map[string]interface{}{
		"total":       total,
		"total_pages": int((total + int64(per_page) - 1) / int64(per_page)),
		"page":        page,
		"per_page":    per_page,
		"data":        data,
	}

	return result
}

func PrettyPrint(v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println("pretty-print error:", err)
	} else {
		fmt.Println(string(b))
	}
}
