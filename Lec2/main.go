package main

import (
	"fmt"
	"log"
	"net/http"
)

//SayHello ...
func SayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world from docker!")
}

//SayGoodbye ...
func SayGoodbye(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Goodbye from docker!")
}

func main() {
	http.HandleFunc("/hello", SayHello)
	http.HandleFunc("/bye", SayGoodbye)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
