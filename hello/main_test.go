package main

import (
  "net/http"
  "net/http/httptest"
  "testing"

  "github.com/stretchr/testify/assert"
)

func TestHealthzRoute(t *testing.T) {
  r := setupRouter()

  w := httptest.NewRecorder()
  req, _ := http.NewRequest("GET", "/healthz", nil)
  r.ServeHTTP(w, req)

  assert.Equal(t, http.StatusOK, w.Code)
  assert.Equal(t, "{\"health\":\"ok\"}", w.Body.String())
}

func TestHelloRoute(t *testing.T) {
  r := setupRouter()

  w := httptest.NewRecorder()
  req, _ := http.NewRequest("GET", "/hello/foobar", nil)
  r.ServeHTTP(w, req)

  assert.Equal(t, http.StatusOK, w.Code)
  assert.Equal(t, "Hello foobar", w.Body.String())
}

func TestMetricsRoute(t *testing.T) {
  r := setupRouter()

  w := httptest.NewRecorder()
  req, _ := http.NewRequest("GET", "/metrics", nil)
  r.ServeHTTP(w, req)

  assert.Equal(t, http.StatusOK, w.Code)
}
