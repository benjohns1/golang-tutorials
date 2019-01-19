package main

import (
	"fmt"
	"net/http"
)

type myMux struct{}

func (p *myMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayHello(w, r)
		return
	}
	http.NotFound(w, r)
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello myroute!")
}

func main() {
	mux := &myMux{}
	http.ListenAndServe(":9090", mux)
}
