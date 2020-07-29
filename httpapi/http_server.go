package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/rizaramadan/go101/web"
)

func home(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "this is home, %q \n", html.EscapeString(r.URL.Path))
	if err != nil {
		log.Println(err)
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "hello, %q \n", html.EscapeString(r.URL.Path))
		if err != nil {
			log.Println(err)
		}
	})

	web.NewRoute("/home", home)

	log.Println("listen localhost 8989")
	log.Fatal(http.ListenAndServe(":8989", nil))
}
