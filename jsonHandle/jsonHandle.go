package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type helloWorldResponse struct {
	Message string
}

func main() {
	port := 8080
	http.HandleFunc("/", helloWorldJsonHandler)
	log.Printf("Starting server on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
func helloWorldJsonHandler(w http.ResponseWriter, r *http.Request) {
	response := helloWorldResponse{Message: "Hello World"}
	data, err := json.Marshal(response)
	if err != nil {
		panic("Oops something went wrong")
	}
	fmt.Fprintln(w, string(data))
}
