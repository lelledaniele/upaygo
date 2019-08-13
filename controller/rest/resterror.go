package apprest

import (
	"encoding/json"
	"net/http"
)

type RESTError struct {
	M string `json:"error"`
}

func (e *RESTError) Error() string {
	return e.M
}

// Write the error in the http response
func handleError(w *http.ResponseWriter, s string) {
	e := RESTError{M: s}
	_ = json.NewEncoder(*w).Encode(e)
}
