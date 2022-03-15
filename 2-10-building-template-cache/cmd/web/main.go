package main

import (
	"fmt"
	"net/http"

	"github.com/Riyaz-khan-shuvo/go-practice/2-10-building-template-cache/pkg/handlers"
)

const portNumber string = ":8500"

func main() {
	fmt.Println("listing the port http://localhost" + portNumber)

	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/about", handlers.AboutPage)
	http.ListenAndServe(portNumber, nil)
}
