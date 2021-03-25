package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"
)

// This function performs checks such as:
// StatusCode
// content-type
// content (body) and value.
func TestTimeZoneCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/timezone", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(TimeZoneHandler)
	handler.ServeHTTP(rr, req)
	t.Run("Verify status code: 200 OK for GET http method", func(t *testing.T) {
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}
	})

	t.Run("Verify content_type for GET http method", func(t *testing.T) {
		if content_type := rr.Result().Header.Get("Content-Type"); content_type != "application/json; charset=utf-8" {
			t.Errorf("BAD content_type : %s", content_type)
		}
	})

	t.Run("Verify Content for GET http method", func(t *testing.T) {
		expected := "Asia/Tokyo"
		var resultRequest TimeZone
		json.NewDecoder(rr.Body).Decode(&resultRequest)
		if resultRequest.TimeZone != expected {
			t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
		}
	})
}

// Test invalide method to "/timezone"
func TestBadMethodTimeZoneHandler(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/timezone", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(TimeZoneHandler)
	handler.ServeHTTP(rr, req)
	t.Run("Verify status code: 200 is not possible for DELETE http method", func(t *testing.T) {
		if status := rr.Code; status == http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}
	})
}

// Verify Default time zone
func TestDefaultTimeZone(t *testing.T) {
	if defaultTimeZone != "Asia/Tokyo" {
		t.Errorf("Bad default TimeZone set : %s", defaultTimeZone)
	}
}