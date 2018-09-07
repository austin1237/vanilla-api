package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/user/api/server"
	"github.com/user/api/stats"
)

func TestStats(t *testing.T) {
	req, err := http.NewRequest("GET", "/stats", nil)
	if err != nil {
		t.Fatal(err)
	}

	sStats := stats.New()
	rr := httptest.NewRecorder()
	handler := Stats(sStats)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"total":0,"average":0}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestShutDown(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	done := make(chan bool, 1)
	mockServer := server.New(done, "3000")
	rr := httptest.NewRecorder()
	handler := ShutDown(mockServer)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `Shutting Down`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestHash(t *testing.T) {
	reader := strings.NewReader("password=angryMonkey")
	req, err := http.NewRequest("POST", "/", reader)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	sStats := stats.New()
	rr := httptest.NewRecorder()
	handler := Hash(sStats)

	// populates the context's startTime
	ctx := req.Context()
	ctx = context.WithValue(ctx, "startTime", time.Now())
	req = req.WithContext(ctx)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `ZEHhWB65gUlzdVwtDQArEyx+KVLzp/aTaRaPlBzYRIFj6vjFdqEb0Q5B8zVKCZ0vKbZPZklJz0Fd7su2A+gf7Q==`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestHash_EmptyPassword(t *testing.T) {
	reader := strings.NewReader("password=")
	req, err := http.NewRequest("POST", "/", reader)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	sStats := stats.New()
	rr := httptest.NewRecorder()
	handler := Hash(sStats)

	// populates the context's startTime
	ctx := req.Context()
	ctx = context.WithValue(ctx, "startTime", time.Now())
	req = req.WithContext(ctx)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	// Check the response body is what we expect.
	expected := "password was not found in form data\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: want %v got %v",
			expected, rr.Body.String())
	}
}

func TestHash_MissingPassword(t *testing.T) {
	reader := strings.NewReader("notpassword=")
	req, err := http.NewRequest("POST", "/", reader)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	sStats := stats.New()
	rr := httptest.NewRecorder()
	handler := Hash(sStats)

	// populates the context's startTime
	ctx := req.Context()
	ctx = context.WithValue(ctx, "startTime", time.Now())
	req = req.WithContext(ctx)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	// Check the response body is what we expect.
	expected := "password was not found in form data\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: want %v got %v",
			expected, rr.Body.String())
	}
}
