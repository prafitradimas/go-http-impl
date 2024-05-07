package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Validator interface {
	IsValid(ctx context.Context) (problems map[string]string)
}

func RequestIsValid[T Validator](request *http.Request) (T, map[string]string, error) {
	var val T

	if err := json.NewDecoder(request.Body).Decode(&val); err != nil {
		return val, nil, fmt.Errorf("decode json: %w", err)
	}

	if problems := val.IsValid(request.Context()); len(problems) > 0 {
		return val, problems, fmt.Errorf("invalid %T: %d problems", val, len(problems))
	}

	return val, nil, nil
}
