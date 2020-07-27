// This sample code implement a simple web service.
package main

import (
	"log"
	"net/http"

	"go_inaction/chapter9/listing17/handlers"
)

// main is the entry point for the application.
func main() {
	handlers.Routes()

	log.Println("listener : Started : Listening on :4000")
	http.ListenAndServe(":4000", nil)
}
