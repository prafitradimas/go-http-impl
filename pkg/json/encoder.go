package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// encode server's response with status and response body as params
func Encode[T any](writter http.ResponseWriter, _ *http.Request, status int, value T) error {
	writter.Header().Set("Content-Type", "application/json")
	writter.WriteHeader(status)

	if err := json.NewEncoder(writter).Encode(&value); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}

	return nil
}

// decode incoming request then returns request.Body or Error
func Decode[T any](request *http.Request) (T, error) {
	var val T
	if err := json.NewDecoder(request.Body).Decode(&val); err != nil {
		return val, fmt.Errorf("decode json: %w", err)
	}

	return val, nil
}
