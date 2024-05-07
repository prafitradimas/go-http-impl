package middlewares

import (
	"net/http"
	"strings"

	"github.com/prafitradimas/go-http-impl/internal/types"
	json "github.com/prafitradimas/go-http-impl/pkg/json"
)

func AdminOnly(handler http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(writter http.ResponseWriter, request *http.Request) {
		_, found := strings.CutPrefix(request.Header.Get("Authorization"), "Bearer ")
		if !found {
			json.Encode[any](writter, request, http.StatusUnauthorized, &types.ApiResponse[any]{
				Success:    false,
				StatusCode: http.StatusUnauthorized,
				Message:    "Request is unauthorized",
			})
			return
		}

		handler(writter, request)
	})
}
