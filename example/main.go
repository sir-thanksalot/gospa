package main

import (
	"net/http"

	spa "github.com/sir-thanksalot/gospa"
)

func main() {
	spaHandler := spa.NewSpaHandler("example/static", "index.html")
	http.Handle("/", spaHandler)
	http.ListenAndServe(":5000", nil)
}
