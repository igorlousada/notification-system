package main

import (
	"testing"
	"os"
	"net/http"
	"net/http/httptest"
	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
)


func TestPostEndpoint(t *testing.T) {
	os.Setenv("test_env", "true")
	app := setupRoutes()

	req := httptest.NewRequest(http.MethodPost, "/send-notification", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}