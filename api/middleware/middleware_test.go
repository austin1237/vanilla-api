package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestStartTime(t *testing.T) {
	req, err := http.NewRequest("POST", "/hash", nil)
	if err != nil {
		t.Fatal(err)
	}

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if val, ok := r.Context().Value("startTime").(time.Time); !ok {
			t.Errorf("startTime not in request context: got %q", val)
		}
	})

	rr := httptest.NewRecorder()
	// func StartTime(h http.Handler) http.Handler
	// Stores an "startTime" in the request context.
	handler := StartTime(testHandler)
	handler.ServeHTTP(rr, req)
}

func TestGetOnly_Get(t *testing.T) {
	req, err := http.NewRequest("GET", "/hash", nil)
	if err != nil {
		t.Fatal(err)
	}

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	rr := httptest.NewRecorder()
	// func GetOnly(h http.Handler) http.Handler
	// Returns a 404 on any non GET Request.
	handler := GetOnly(testHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestGetOnly_Post(t *testing.T) {
	req, err := http.NewRequest("POST", "/hash", nil)
	if err != nil {
		t.Fatal(err)
	}

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	rr := httptest.NewRecorder()
	// func GetOnly(h http.Handler) http.Handler
	// Returns a 404 on any non GET Request.
	handler := GetOnly(testHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}

func TestGetOnly_Put(t *testing.T) {
	req, err := http.NewRequest("PUT", "/hash", nil)
	if err != nil {
		t.Fatal(err)
	}

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	rr := httptest.NewRecorder()
	// func GetOnly(h http.Handler) http.Handler
	// Returns a 404 on any non GET Request.
	handler := GetOnly(testHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}

func TestGetOnly_Delete(t *testing.T) {
	req, err := http.NewRequest("Delete", "/hash", nil)
	if err != nil {
		t.Fatal(err)
	}

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	rr := httptest.NewRecorder()
	// func GetOnly(h http.Handler) http.Handler
	// Returns a 404 on any non GET Request.
	handler := GetOnly(testHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}

func TestPostOnly_Get(t *testing.T) {
	req, err := http.NewRequest("GET", "/hash", nil)
	if err != nil {
		t.Fatal(err)
	}

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	rr := httptest.NewRecorder()
	// func PostOnly(h http.Handler) http.Handler
	// Returns a 404 on any non Post Request.
	handler := PostOnly(testHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}

func TestPostOnly_Post(t *testing.T) {
	req, err := http.NewRequest("POST", "/hash", nil)
	if err != nil {
		t.Fatal(err)
	}

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	rr := httptest.NewRecorder()
	// func PostOnly(h http.Handler) http.Handler
	// Returns a 404 on any non Post Request.
	handler := PostOnly(testHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestPostOnly_Put(t *testing.T) {
	req, err := http.NewRequest("PUT", "/hash", nil)
	if err != nil {
		t.Fatal(err)
	}

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	rr := httptest.NewRecorder()
	// func PostOnly(h http.Handler) http.Handler
	// Returns a 404 on any non Post Request.
	handler := PostOnly(testHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}

func TestPostOnly_DELETE(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/hash", nil)
	if err != nil {
		t.Fatal(err)
	}

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	rr := httptest.NewRecorder()
	// func PostOnly(h http.Handler) http.Handler
	// Returns a 404 on any non Post Request.
	handler := PostOnly(testHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}
