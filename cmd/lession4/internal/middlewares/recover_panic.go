package middlewares

import (
	"net/http"
)

func RecoverPanicMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// implement logic of middleware
		defer func() {
			if err := recover(); err != nil {
				// log
				// close connection
				w.Header().Set("Connection", "close")

				http.Error(w, "err", 500)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
