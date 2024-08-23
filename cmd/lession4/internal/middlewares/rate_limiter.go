package middlewares

import (
	"net/http"

	"golang.org/x/time/rate"
)

func RateLimiterMiddleware(next http.Handler) http.Handler {
	limiter := rate.NewLimiter(2, 4)


	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request)  {
		// implement logic of middleware
		if !limiter.Allow() {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
		}

		next.ServeHTTP(w, r)
	})
}