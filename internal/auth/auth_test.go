package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	apiKey := "test-api-key"
	normalHeader := http.Header{}

	normalHeader.Add("Authorization", "ApiKey "+apiKey)

	key, err := GetAPIKey(normalHeader)

	if err != nil {
		t.Errorf("Error getting normal key")
		return
	}

	if key != apiKey {
		t.Errorf("Keys do not match")
	}
}

func TestGetAPIKeyWithError(t *testing.T) {
	badHeader := http.Header{}

	_, err := GetAPIKey(badHeader)

	if err == nil {
		t.Errorf("No error received on invalid header")
	}
}
