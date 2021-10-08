package exception

import "net/http"

func Throw(w http.ResponseWriter, statusCode int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(statusCode)

	var message string

	if msg == "" {
		message = http.StatusText(statusCode)
	} else {
		message = msg
	}

	w.Write([]byte(message))
}

func InternalServerError(w http.ResponseWriter, msg string) {
	Throw(w, http.StatusInternalServerError, msg)
}
