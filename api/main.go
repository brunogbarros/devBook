package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Carregar()
	r := router.GetRouter()

	fmt.Println("Server running on localhost: ", config.Port)
	log.Fatal(http.ListenAndServe(":5000", r))

}
