package httpserver

import (
	"fmt"
	"net/http"
)

const robotsTXT string = `
User-Agent: *

Disallow: /harming/humans
Disallow: /ignoring/human/orders
Disallow: /harm/to/self

Disallow: /
`

//NoRobotsHandler default robots.txt handler. Disallow all.
func NoRobotsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, robotsTXT)
}
