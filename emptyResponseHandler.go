package httpserver

import (
	"fmt"
	"net/http"
)

//EmptyReponseHandler stub handler, serves empty response.
func EmptyReponseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "")
}
