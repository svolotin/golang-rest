// Copyright (c) 2023 ...

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheck_IT(t *testing.T) {
	router := setupRouter()

	responseRecorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/healthcheck", nil)
	router.ServeHTTP(responseRecorder, request)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Equal(t, "OK", responseRecorder.Body.String())
}
