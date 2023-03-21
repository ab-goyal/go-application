package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
)

func TestFormHandler(t *testing.T) {
	// Create a new request with form data.
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

func TestPrintKVMap(t *testing.T) {
	// Create a new buffer to capture the output
	buf := new(bytes.Buffer)

	// Create a sample key-value map
	kvMap := map[string]string{
		"name": "John",
		"age":  "30",
		"city": "New York",
	}

	// Call the function with the sample map and the buffer
	printKVMap(kvMap, fakeResponseWriter)

	// Check the output against the expected result
	expected := "name: John\nage: 30\ncity: New York\n"
	if buf.String() != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, buf.String())
	}

	type fakeResponseWriter struct {
		buf *bytes.Buffer
	}
	
	func (w fakeResponseWriter) Header() http.Header {
		return http.Header{}
	}
	
	func (w fakeResponseWriter) Write(b []byte) (int, error) {
		return w.buf.Write(b)
	}
	
	func (w fakeResponseWriter) WriteHeader(statusCode int) {
	}
}
