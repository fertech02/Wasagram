package api

import (
	"net/http"
	"strings"
)

func CheckValidAuth(r *http.Request) bool {

	// Get the Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return false
	}

	authParts := strings.Fields(authHeader)
	if len(authParts) != 2 || authParts[0] != BEAR || authParts[1] == "null" {
		return false
	}

	return true
}

// IsAuthorized checks if the user is authorized
func CheckIdAuthorized(r *http.Request, id string) int {

	// Get the Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return 1
	}

	authHeaderParts := strings.Fields(authHeader)
	if len(authHeaderParts) != 2 || authHeaderParts[0] != BEAR || authHeaderParts[1] == "null" {
		return 1
	}

	token := authHeaderParts[1]
	if token != id {
		return 2
	}

	return 0
}

func GetIdFromBearer(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")

	authHeaderParts := strings.Fields(authHeader)
	token := authHeaderParts[1]

	return token
}
