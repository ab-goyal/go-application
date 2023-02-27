package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTime(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:3000/mytime", nil)
	if err != nil {
		t.Fatalf("could not created request: %v", err)
	}
	rec := httptest.NewRecorder()
	timeHandler(rec, req)

	res := rec.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.Status)
	}
}
