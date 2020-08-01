package main

import (
	"fmt"
	"github.com/rizaramadan/go101/web"
	"html"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "this is home, %q \n",
		html.EscapeString(r.URL.Path))
	if err != nil {
		log.Println(err)
	}
}

func main() {
	route := web.NewRoute("/", home)
	route.AddRoute()

	log.Println("listen localhost 8989")
	log.Fatal(http.ListenAndServe(":8989", nil))
}
