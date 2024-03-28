package main

import (
	web_ser "Web_Art/LocalServer"
	"fmt"
	"log"
	"net/http"
)

func main() {

	// Start the server
	http.HandleFunc("/", web_ser.HandlerFunc)

	fmt.Println("Visit the following website:\nhttp://localhost:8080")

	fmt.Println("Server listening on port 8080...")

	// Close the server if there is any problem while runing
	log.Fatal(http.ListenAndServe(":8080", nil))
}
