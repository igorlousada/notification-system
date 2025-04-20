package main

import (
	"testing"
	"bytes"
	"os"
	"net/http"
	"net/http/httptest"
	"encoding/json"
	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
)

func setup() *fiber.App{
	os.Setenv("test_env", "true")
	app := setupRoutes()
	return app
}

func TestPostEndpoint(t *testing.T) {
	app := setup()
	requestBody := map[string]string{"user_uuid": "123", "message": "purchase alert - $45"}
	parsedBody, _ := json.Marshal(requestBody)

	req := httptest.NewRequest(http.MethodPost, "/send-notification", bytes.NewReader(parsedBody))
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}

func TestPostInvalidJsonBody (t *testing.T) {
	requestBody := ` {
		"user_uuid: "123",
		"message": "sample message"
	}
	`
	app := setup()
	req := httptest.NewRequest(http.MethodPost, "/send-notification", bytes.NewReader([]byte(requestBody)))
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
}

func TestPostEmptyUserUuid (t *testing.T) {
	requestBody := ` {
		"user_uuid": "",
		"message": "sample message"
	}
	`
	app := setup()
	req := httptest.NewRequest(http.MethodPost, "/send-notification", bytes.NewReader([]byte(requestBody)))
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusUnprocessableEntity, resp.StatusCode)
}

func TestPostEmptyMessage (t *testing.T) {
	requestBody := ` {
		"user_uuid": "123",
		"message": ""
	}
	`
	app := setup()
	req := httptest.NewRequest(http.MethodPost, "/send-notification", bytes.NewReader([]byte(requestBody)))
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusUnprocessableEntity, resp.StatusCode)
}


