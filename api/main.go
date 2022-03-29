package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := router.GetRouter()

	fmt.Println("Server running on localhost:5000")
	log.Fatal(http.ListenAndServe(":5000", r))

}
