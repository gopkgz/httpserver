package httpserver

import (
	"fmt"
	"net/http"
)

//EmptyHandler stub handler, serves empty response.
func EmptyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "")
}
