package main

import (
	"fmt"

	"net/http"

	routing "learn/routers"
	// "github.com/go-chi/chi"
)

const HOST string = "127.0.0.1"
const PORT string = "3333"

func main() {
	r := routing.Routing()

	//  meddleware in the routing package

	fmt.Println("\033[32m+ Starting server on : \033[37m", HOST+":"+PORT)
	http.ListenAndServe(HOST+":"+PORT, r)

}
