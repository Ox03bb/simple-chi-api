package main

import (
	"fmt"

	"net/http"

	models "learn/models"
	routing "learn/routers"
	// "github.com/go-chi/chi"
)

const HOST string = "127.0.0.1"
const PORT string = "3333"

func main() {
	_, err := models.InitDB()
	if err != nil {
		fmt.Println("\033[31m- Error connecting to the database: \033[37m", err)
		return
	}

	r := routing.Routing()

	//  meddleware in the routing package

	fmt.Println("\033[32m+ Starting server on : \033[37m", HOST+":"+PORT)
	http.ListenAndServe(HOST+":"+PORT, r)

}
