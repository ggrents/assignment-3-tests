package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/highxshell/greenlight/internal/data"
)

type mockUsersModel struct{}

func (m *mockUsersModel) Insert(user *data.User) error {
	return nil
}

func (m *mockUsersModel) GetForToken(scope, tokenPlaintext string) (*data.User, error) {
	return &data.User{}, nil
}

func (m *mockUsersModel) Update(user *data.User) error {
	return nil
}

func (m *mockTokensModel) DeleteAllForUser(scope string, userID int64) error {
	return nil
}

func TestRegisterUserHandler(t *testing.T) {
	app := &application{
		models: &models{
			Users: &mockUsersModel{},
		},
	}

	jsonBody := `{"name": "testuser", "email": "test@example.com", "password": "testpassword"}`
	req := httptest.NewRequest("POST", "/v1/users", bytes.NewBufferString(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	app.registerUserHandler(rr, req)

	if rr.Code != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusAccepted)
	}
}

func TestActivateUserHandler(t *testing.T) {
	app := &application{
		models: &models{
			Users: &mockUsersModel{},
		},
	}

	jsonBody := `{"token": "testtoken"}`
	req := httptest.NewRequest("POST", "/v1/tokens/authentication", bytes.NewBufferString(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	app.activateUserHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
	}
}

func TestUpdateUserPasswordHandler(t *testing.T) {
	app := &application{
		models: &models{
			Users: &mockUsersModel{},
		},
	}

	jsonBody := `{"password": "newpassword", "token": "testtoken"}`
	req := httptest.NewRequest("POST", "/v1/users/password", bytes.NewBufferString(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	app.updateUserPasswordHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
	}
}
