// Copyright (c) 2023 ...

package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func TestHealthCheck(t *testing.T) {
	response := httptest.NewRecorder()
	testCtx, _ := gin.CreateTestContext(response)

	// Invoke controller function
	HealthCheck(testCtx)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "OK", response.Body.String())
}
