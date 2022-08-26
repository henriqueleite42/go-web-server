package main

import (
	"fmt"
	"log"
	"net/http"

	"go-web-server/routes"
)


func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)

	http.HandleFunc("/form", routes.FormHandler)

	http.HandleFunc("/hello", routes.HelloHandler)

	fmt.Printf("Starting server at port 8080\n")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}