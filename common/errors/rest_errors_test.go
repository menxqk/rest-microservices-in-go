package errors

import (
	"errors"
	"net/http"
	"testing"
)

func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError("this is the message")

	if err == nil {
		t.Fatal("err should not be nil")
	}
	if err.Status != http.StatusBadRequest {
		t.Errorf("expected status code %d, got %d", http.StatusBadRequest, err.Status)
	}
	if err.Message != "this is the message" {
		t.Errorf("wrong error message: %s", err.Message)
	}
	if err.Error != "bad_request" {
		t.Errorf("wrong err.Error: %s", err.Error)
	}
}

func TestNewUnauthorizedError(t *testing.T) {
	err := NewUnauthorizedError("this is the message")

	if err == nil {
		t.Fatal("err should not be nil")
	}
	if err.Status != http.StatusUnauthorized {
		t.Errorf("expected status code %d, got %d", http.StatusUnauthorized, err.Status)
	}
	if err.Message != "this is the message" {
		t.Errorf("wrong error message: %s", err.Message)
	}
	if err.Error != "unauthorized" {
		t.Errorf("wrong err.Error: %s", err.Error)
	}
}

func TestNewNotFoundError(t *testing.T) {
	err := NewNotFoundError("this is the message")

	if err == nil {
		t.Fatal("err should not be nil")
	}
	if err.Status != http.StatusNotFound {
		t.Errorf("expected status code %d, got %d", http.StatusNotFound, err.Status)
	}
	if err.Message != "this is the message" {
		t.Errorf("wrong error message: %s", err.Message)
	}
	if err.Error != "not_found" {
		t.Errorf("wrong err.Error: %s", err.Error)
	}
}

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("this is the message", errors.New("database error"))

	if err == nil {
		t.Fatal("err should not be nil")
	}
	if err.Status != http.StatusInternalServerError {
		t.Errorf("expected status code %d, got %d", http.StatusInternalServerError, err.Status)
	}
	if err.Message != "this is the message" {
		t.Errorf("wrong error message: %s", err.Message)
	}
	if err.Error != "internal_server_error" {
		t.Errorf("wrong err.Error: %s", err.Error)
	}
	if err.Causes == nil {
		t.Fatal("causes should not be nil")
	}
	if len(err.Causes) != 1 {
		t.Errorf("expected length of causes equal to 1, got %d", len(err.Causes))
	}
	if err.Causes[0] != "database error" {
		t.Errorf("wrong error cause: %s", err.Causes[0].(string))
	}
}
