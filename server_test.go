package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

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
	// req, err := http.NewRequest("GET", "localhost:3000/mytime", nil)
	// if err != nil {
	// 	t.Fatalf("could not create request: %v", err)
	// }
	// rec := httptest.NewRecorder()
	// timeHandler(rec, req)
	// resp := rec.Result()
	// body := rec.Body
	// fmt.Println(resp.StatusCode)
	// fmt.Println(resp.Header)
	// fmt.Println(body)
	// res := rec.Result()
	// if res.StatusCode != http.StatusOK {
	// 	t.Errorf("expected status OK; got %v", res.Status)
	// }
	//

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
