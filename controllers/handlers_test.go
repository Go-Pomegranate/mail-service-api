package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var controller = Controller{}

func TestHealthCheckHandler(t *testing.T) {
	//Create a request to pass to our handler
	req, err := http.NewRequest("GET", "/index", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.Index)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}
