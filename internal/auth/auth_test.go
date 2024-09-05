package auth

import (
	"net/http"
	"testing"
)

// TestGetApiKey gets the ApiKey from the Authorization header
func TestGetApiKey(t *testing.T) {
	expected := "theapikey"
	headers := http.Header{
		"Authorization": []string{"ApiKey " + expected},
	}

	actual, err := GetAPIKey(headers)

	if err != nil {
		t.Fatalf("GetAPIKey returned an error: %v", err)
	}

	if actual != expected {
		t.Fatalf("GetAPIKey returned the wrong value: %v", actual)
	}
}
