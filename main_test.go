package main

import (
  "net/http"
  "net/http/httptest"
  "testing"
)

var expectedPayload = `{"myapplication":[{"version":"1.0","description":"pre-interview technical test","lastcommitsha":"abc123"}]}`

func TestHealthCheckHandler(t *testing.T) {
  // Create a request to healthcheck endpoint
  req, err := http.NewRequest("GET", "/health-check", nil)
  if err != nil {
      t.Fatal(err)
  }

  // Set ResponseRecorder to record the response.
  rr := httptest.NewRecorder()
  handler := http.HandlerFunc(healthcheckHandler)

  handler.ServeHTTP(rr, req)

  // Check the status code is what we expect.
  if status := rr.Code; status != http.StatusOK {
    t.Errorf("handler returned wrong status code: got %v want %v",
      status, http.StatusOK)
  }

  // Check the response body is what we expect.
  if rr.Body.String() != expectedPayload {
    t.Errorf(
      `"handler returned unexpected body: got:
      %v
      want:
      %v"`,
      rr.Body.String(), expectedPayload)
  }
}
