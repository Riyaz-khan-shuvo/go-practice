package main

import (
	"fmt"
	"net/http"

	"github.com/Riyaz-khan-shuvo/go-practice/2-13-optimizing-the-template/pkg/config"
	"github.com/Riyaz-khan-shuvo/go-practice/2-13-optimizing-the-template/pkg/handlers"
	"github.com/Riyaz-khan-shuvo/go-practice/2-13-optimizing-the-template/pkg/render"
)

const portNumber string = ":5559"

func main() {

	var app config.AppConfig
	tc, err := render.CreateTemplate()
	if err != nil {
		fmt.Println(err)
	}
	app.TemplateCache = tc

	render.NewTemplate(&app)

	fmt.Println("Listing the port http://localhost" + portNumber)
	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/about", handlers.AboutPage)
	http.ListenAndServe(portNumber, nil)

}
