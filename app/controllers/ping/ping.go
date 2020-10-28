package ping

import "net/http"

// Ping handler
func Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pong!"))
}
