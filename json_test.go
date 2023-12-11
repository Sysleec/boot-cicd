package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRespondWithJSON(t *testing.T) {
	w := httptest.NewRecorder()

	payload := map[string]string{
		"name": "John",
	}

	respondWithJSON(w, http.StatusOK, payload)

	var expectedBody = []byte(`{"name":"John"}`)
	if !bytes.Equal(w.Body.Bytes(), expectedBody) {
		t.Errorf("Expected response body to be %s, got %s", expectedBody, w.Body.Bytes())
	}

	if w.Code != http.StatusOK {
		t.Errorf("Expected response code to be %d, got %d", http.StatusOK, w.Code)
	}
}

func TestRespondWithError(t *testing.T) {
	w := httptest.NewRecorder()

	msg := "Internal server error"

	respondWithError(w, http.StatusInternalServerError, msg)

	var expectedBody = []byte(`{"error":"Internal server error"}`)
	if !bytes.Equal(w.Body.Bytes(), expectedBody) {
		t.Errorf("Expected response body to be %s, got %s", expectedBody, w.Body.Bytes())
	}

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected response code to be %d, got %d", http.StatusInternalServerError, w.Code)
	}
}
