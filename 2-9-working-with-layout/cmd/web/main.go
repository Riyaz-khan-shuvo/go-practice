package main

import (
	"fmt"
	"net/http"

	"github.com/Riyaz-khan-shuvo/go-practice/2-9-working-with-layout/pkg/handlers"
)

const portNumber string = ":9000"

func main() {
	fmt.Println("listing the port http://localhost" + portNumber)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	http.ListenAndServe(portNumber, nil)
}
