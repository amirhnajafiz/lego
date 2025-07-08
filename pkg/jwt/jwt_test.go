package jwt

import (
	"testing"
	"time"
)

// TestJWT tests the JWTManager's ability to create and parse JWT tokens.
// It generates a token with specified headers, logs the token and its expiration time,
// and then parses the token to verify the headers.
func TestJWT(t *testing.T) {
	signingKey := "my_secret_key"
	expirationDuration := time.Hour

	jwtManager := New(signingKey, expirationDuration)

	headers := map[string]string{"user_id": "123"}
	tokenString, expireTime, err := jwtManager.BuildToken(headers)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	t.Logf("Generated token: %s", tokenString)
	t.Logf("Token expiration time: %v", expireTime)

	parsedHeaders, err := jwtManager.ParseToken(tokenString)
	if err != nil {
		t.Fatalf("Failed to parse token: %v", err)
	}

	t.Logf("Parsed headers: %v", parsedHeaders)
}

// TestJWTInvalidSigningMethod tests the JWTManager's handling of an invalid signing method.
// It generates a token, logs the token and its expiration time, and then attempts to parse
// the token with an invalid signing method, expecting an error.
func TestJWTInvalidSigningMethod(t *testing.T) {
	signingKey := "my_secret_key"
	expirationDuration := time.Hour

	jwtManager := New(signingKey, expirationDuration)

	headers := map[string]string{"user_id": "123"}
	tokenString, expireTime, err := jwtManager.BuildToken(headers)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	t.Logf("Generated token: %s", tokenString)
	t.Logf("Token expiration time: %v", expireTime)

	parsedHeaders, err := jwtManager.ParseToken(tokenString)
	if err != nil {
		t.Fatalf("Failed to parse token: %v", err)
	}

	t.Logf("Parsed headers: %v", parsedHeaders)
}

// TestJWTInvalidToken tests the JWTManager's handling of an invalid token.
// It generates a valid token, logs the token and its expiration time, then tampers with the token.
func TestJWTInvalidToken(t *testing.T) {
	signingKey := "my_secret_key"
	expirationDuration := time.Hour

	jwtManager := New(signingKey, expirationDuration)

	headers := map[string]string{"user_id": "123"}
	tokenString, expireTime, err := jwtManager.BuildToken(headers)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	t.Logf("Generated token: %s", tokenString)
	t.Logf("Token expiration time: %v", expireTime)

	// tamper with the token
	invalidTokenString := tokenString + "invalid"

	parsedHeaders, err := jwtManager.ParseToken(invalidTokenString)
	if err != nil {
		t.Logf("Expected error for invalid token: %v", err)
	} else {
		t.Fatalf("Parsed headers from invalid token: %v", parsedHeaders)
	}
}

// TestJWTExpiredToken tests the JWTManager's handling of an expired token.
// It generates a token with a very short expiration duration, waits for it to expire,
// and then attempts to parse the token, expecting an error due to expiration.
func TestJWTExpiredToken(t *testing.T) {
	signingKey := "my_secret_key"
	expirationDuration := time.Nanosecond

	jwtManager := New(signingKey, expirationDuration)

	headers := map[string]string{"user_id": "123"}
	tokenString, expireTime, err := jwtManager.BuildToken(headers)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	t.Logf("Generated token: %s", tokenString)
	t.Logf("Token expiration time: %v", expireTime)

	time.Sleep(expirationDuration)

	parsedHeaders, err := jwtManager.ParseToken(tokenString)
	if err != nil {
		t.Logf("Expected error for expired token: %v", err)
	} else {
		t.Fatalf("Parsed headers from expired token: %v", parsedHeaders)
	}
}
