package auth

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAPIKey(t *testing.T) {

	headers := http.Header{}
	headers.Set("Authorization", "ApiKey your-api-key")

	// Test case: Valid authorization header
	apiKey, err := GetAPIKey(headers)
	assert.NoError(t, err)
	assert.Equal(t, "your-api-key", apiKey)

	// Test case: No authorization header included
	headers.Del("Authorization")
	_, err = GetAPIKey(headers)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, ErrNoAuthHeaderIncluded))

	// Test case: Malformed authorization header
	headers.Set("Authorization", "InvalidHeader")
	_, err = GetAPIKey(headers)
	assert.Error(t, err)
	assert.EqualError(t, err, "malformed authorization header")
}
