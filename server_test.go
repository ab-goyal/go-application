package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
)

func TestFormHandler(t *testing.T) {
	// Create a new request with form data
	form := url.Values{}
	form.Add("name", "John")
	form.Add("address", "123 Main St")
	req, err := http.NewRequest("POST", "/form", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatal(err)
	}

	// Create a new response recorder
	rr := httptest.NewRecorder()

	// Call the formHandler function with the new request and response recorder
	handler := http.HandlerFunc(formHandler)
	handler.ServeHTTP(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	expected := "POST request successful\nName = John\nAddress = 123 Main St\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestTime(t *testing.T) {
	req, err := http.NewRequest("GET", "/mytime", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(timeHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expectedTime := time.Now().Format("01-02-2006 15:04:05 Monday")
	if body := rr.Body.String(); body != "The current time is: "+expectedTime+"\n"+message {
		t.Errorf("handler returned unexpected body: got %v want %v",
			body, "The current time is: "+expectedTime+"\n"+message)
	}

}

func TestHello(t *testing.T) {
	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		helloHandler(w, r, "Alice")
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expectedResponse := "Hello, Alice!\n" + message
	if body := rr.Body.String(); body != expectedResponse {
		t.Errorf("handler returned unexpected body: got %v want %v",
			body, expectedResponse)
	}
}
