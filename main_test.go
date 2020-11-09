package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestRootRouteHandler(t *testing.T) {
	request, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	requestRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(RootRouteHandler)

	handler.ServeHTTP(requestRecorder, request)

	if status := requestRecorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"healthy":true}` + "\n"

	if requestRecorder.Body.String() != expected {
		spew.Dump(requestRecorder.Body.String(), expected)
		t.Errorf("handler returned unexpected body: got %v want %v",
			requestRecorder.Body.String(), expected)
	}
}
