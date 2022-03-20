package main

import (
	"fmt"
	"net/http"

	"github.com/Riyaz-khan-shuvo/go-practice/2-13-practice-more/pkg/handlers"
)

const portNumber string = ":5959"

func main() {
	fmt.Println("Listing the port http://localhost" + portNumber)

	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/about", handlers.AboutPage)

	http.ListenAndServe(portNumber, nil)
}
