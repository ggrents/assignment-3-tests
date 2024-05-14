package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func newTestApplication(t *testing.T) *application {
	return &application{
		models: &models{
			Movies: &mockMoviesModel{},
		},
	}
}

func TestCreateMovieHandler(t *testing.T) {
	app := newTestApplication(t)

	jsonBody := `{"title": "Test Movie", "year": 2022, "runtime": 120, "genres": ["Action", "Thriller"]}`
	req := httptest.NewRequest("POST", "/v1/movies", bytes.NewBufferString(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	app.createMovieHandler(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusCreated)
	}

	location := rr.Header().Get("Location")
	if location == "" {
		t.Error("handler returned empty Location header")
	}
}

func TestShowMovieHandler(t *testing.T) {
	app := newTestApplication(t)

	req := httptest.NewRequest("GET", "/v1/movies/1", nil)
	rr := httptest.NewRecorder()

	app.showMovieHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
	}
}

func TestUpdateMovieHandler(t *testing.T) {
	app := newTestApplication(t)

	jsonBody := `{"title": "Updated Movie", "year": 2023, "runtime": 130, "genres": ["Action", "Drama"]}`
	req := httptest.NewRequest("PATCH", "/v1/movies/1", bytes.NewBufferString(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	app.updateMovieHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
	}
}
