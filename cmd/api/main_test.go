package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_setupRouter(t *testing.T) {
	tests := []struct {
		name       string
		url        string
		wantedCode int
		wantedBody string
	}{
		{
			"Pong",
			"/ping",
			http.StatusOK,
			`{"message":"pong"}`,
		},
		{
			"Products",
			"/products",
			http.StatusOK,
			`[{
  "sku": "000001",
  "name": "BV Lean leather ankle boots",
  "category": "boots",
  "price": {
    "original": 89000,
    "final": 62300,
    "discount_percentage": "30%",
    "currency": "EUR"
  }
}]`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := setupRouter()
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, tt.url, nil)
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wantedCode, w.Code)
			assert.JSONEq(t, tt.wantedBody, w.Body.String())
		})
	}
}
