package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func MakeGinRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestPing(t *testing.T) {
	router := SetupRouter()
	w := MakeGinRequest(router, "GET", "/ping")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.HeaderMap["Content-Type"], "application/json; charset=utf-8")
}
