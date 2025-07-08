package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// list of errors
var (
	ErrSigningMethod = errors.New("error in signing method")
	ErrInvalidToken  = errors.New("token is invalid")
)

// JWTManager is a component for handling JWT tokens and their claims.
type JWTManager struct {
	signingKey         string
	expirationDuration time.Duration
}

// New builds a new auth struct for handling JWT tokens.
func New(k string, e time.Duration) *JWTManager {
	return &JWTManager{
		signingKey:         k,
		expirationDuration: e,
	}
}

// BuildToken creates a new JWT token with the provided headers and expiration duration.
// It returns the token string and its expiration time.
// If an error occurs during token generation, it returns an empty string and the error.
func (j *JWTManager) BuildToken(headers map[string]string) (string, time.Time, error) {
	// create a new token
	token := jwt.New(jwt.SigningMethodHS256)
	expireTime := time.Now().Add(j.expirationDuration)

	// create claims
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = expireTime.Unix()

	// add headers to claims
	for key, value := range headers {
		claims[key] = value
	}

	// generate token string
	tokenString, err := token.SignedString([]byte(j.signingKey))
	if err != nil {
		return "", expireTime, err
	}

	return tokenString, expireTime, nil
}

// ParseToken gets a token string and extracts the data.
// It returns a map of headers if the token is valid, or an error if the token is invalid or parsing fails.
func (j *JWTManager) ParseToken(tokenString string) (map[string]string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", ErrSigningMethod
		}

		return []byte(j.signingKey), nil
	})
	if err != nil {
		return nil, err
	}

	// taking out claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		headers := make(map[string]string)
		for key, value := range claims {
			if str, ok := value.(string); ok {
				headers[key] = str
			}
		}

		return headers, nil
	}

	return nil, ErrInvalidToken
}
