package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"virtustream/controllers"

	"github.com/gorilla/mux"
)

func TestPostFibSequenceNumsWithServer(t *testing.T) {
	r := mux.NewRouter()
	userJSON := `{"number": 1}`
	reader := strings.NewReader(userJSON)
	r.HandleFunc("/fib", controllers.PostFibSequenceNums).Methods("POST")
	server := httptest.NewServer(r)
	defer server.Close()
	fibURL := fmt.Sprintf("%s/fib", server.URL)
	request, err := http.NewRequest("POST", fibURL, reader)
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("HTTP Status expected: 200, got: %d", res.StatusCode)
	}
}
