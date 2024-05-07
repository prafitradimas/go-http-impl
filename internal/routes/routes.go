package routes

import (
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/prafitradimas/go-http-impl/internal/middlewares"
	"github.com/prafitradimas/go-http-impl/internal/types"
	json "github.com/prafitradimas/go-http-impl/pkg/json"
)

// responsible for all the top-level HTTP that applies to all endpoints
//
// ex: CORS, middleware and logging
func SetupHandlers(
	logger *slog.Logger,
	db *sql.DB,
) http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/", http.NotFoundHandler())
	mux.Handle("/admin", middlewares.AdminOnly(func(w http.ResponseWriter, r *http.Request) {}))
	mux.HandleFunc("/greetings", handleGreetings)

	var handler http.Handler = mux
	// add global middlewares
	// handler = someMiddleware(handler)
	// handler = someMiddleware2(handler)
	// handler = someMiddleware3(handler)
	return handler
}

func handleGreetings(w http.ResponseWriter, r *http.Request) {
	json.Encode(w, r, 200, &types.ApiResponse[string]{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "Success",
		Result:     "Greetings!",
	})
}
