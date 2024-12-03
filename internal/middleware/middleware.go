package middleware

import "net/http"

// CORS sets the appropriate headers to allow cross-origin requests
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")                            // Allow all origins, adjust as needed
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")          // Specify allowed methods
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization") // Specify allowed headers

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
