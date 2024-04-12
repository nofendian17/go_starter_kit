package healthcheck

import "net/http"

func (h *handler) Ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("pong"))
		if err != nil {
			return
		}
	}
}
