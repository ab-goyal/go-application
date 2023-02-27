package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTime(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:3000/mytime", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}
	rec := httptest.NewRecorder()
	timeHandler(rec, req)
	resp := rec.Result()
	body := rec.Body
	fmt.Println(resp.StatusCode)
	//fmt.Println(resp.Header)
	fmt.Println(body)
	// res := rec.Result()
	// if res.StatusCode != http.StatusOK {
	// 	t.Errorf("expected status OK; got %v", res.Status)
	// }
	//

}
