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
	if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" || authParts[1] == "null" {
		return false
	}

	return true
}

// IsAuthorized checks if the user is authorized
func CheckAuthorizedId(r *http.Request, id string) (bool, error) {

	// Get the Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return false, nil
	}

	authHeaderParts := strings.Fields(authHeader)
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		return false, nil
	}

	token := authHeaderParts[1]
	if token != id {
		return false, nil
	}

	return true, nil
}

func GetIdFromBearer(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}

	authHeaderParts := strings.Fields(authHeader)
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		return ""
	}

	return authHeaderParts[1]
}
