package main

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestHelloHandler tests the Hello handler function
func TestHelloHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Hello)

	handler.ServeHTTP(rr, req)

	expected := "hello world\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

// TestMethodMiddleware tests the Method middleware
func TestMethodMiddleware(t *testing.T) {
	req, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := Method("GET")(Hello)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	expected := http.StatusText(http.StatusBadRequest) + "\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

// TestLoggingMiddleware tests the Logging middleware
func TestLoggingMiddleware(t *testing.T) {
	// Create a buffer to capture log output
	var logBuffer bytes.Buffer
	originalLogOutput := log.Writer() // Save the original log output
	log.SetOutput(&logBuffer)
	defer log.SetOutput(originalLogOutput) // Restore original log output after the test

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := Logging()(Hello)

	handler.ServeHTTP(rr, req)

	expected := "hello world\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}

	// Check if the log contains the expected output
	logOutput := logBuffer.String()
	if !strings.Contains(logOutput, "/") {
		t.Errorf("log output does not contain expected values: got %v", logOutput)
	}
}

// TestChain tests the Chain function with multiple middlewares
func TestChain(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := Chain(Hello, Method("GET"), Logging())

	handler.ServeHTTP(rr, req)

	expected := "hello world\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
