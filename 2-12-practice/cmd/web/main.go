package main

import (
	"fmt"
	"net/http"

	"github.com/Riyaz-khan-shuvo/go-practice/2-12-practice/pkg/handlers"
)

const portNumber string = ":5555"

func main() {
	fmt.Println("listing the port http://localhost" + portNumber)
	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/about", handlers.AboutPage)
	http.ListenAndServe(portNumber, nil)
}
