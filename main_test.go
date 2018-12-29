package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestRouter(t *testing.T) {
	// Build our expected body

	var body bucketsjson
	body.Type = "bucket"
	body.List = append(body.List, "canguro")
	body.List = append(body.List, "ardilla")

	// Grab our router
	router := SetupRouter()
	// Perform a GET request with that handler.
	w := performRequest(router, "GET", "/test")
	// Assert we encoded correctly,
	// the request gives a 200
	assert.Equal(t, http.StatusOK, w.Code)
	// Convert the Body response to a struct
	var response bucketsjson
	json.Unmarshal([]byte(w.Body.String()), &response)

	assert.Equal(t, body, response)
}
