package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

var Validate = validator.New()

func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("request body is empty")
	}

	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJSON(w http.ResponseWriter, status int, payload any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	response := map[string]any{
		"success": true,
		"data":    payload,
		"error":   nil,
	}

	return json.NewEncoder(w).Encode(response)
}

func WriteError(w http.ResponseWriter, status int, err error) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	response := map[string]any{
		"success": false,
		"error":   err.Error(),
		"data":    nil,
	}

	return json.NewEncoder(w).Encode(response)
}

func GetURLParams(r *http.Request, key string) (*int, error) {
	vars := mux.Vars(r)
	str, ok := vars[key]
	if !ok {
		return nil, fmt.Errorf("missing params %s", key)
	}

	paramKey, err := strconv.Atoi(str)
	if err != nil {
		return nil, fmt.Errorf("invalid params %s", key)
	}

	return &paramKey, nil
}

func GetTokenFromRequest(r *http.Request) string {
	tokenAuth := r.Header.Get("Authorization")
	tokenQuery := r.URL.Query().Get("token")

	if tokenAuth != "" {
		return tokenAuth
	}

	if tokenQuery != "" {
		return tokenQuery
	}

	return ""
}
