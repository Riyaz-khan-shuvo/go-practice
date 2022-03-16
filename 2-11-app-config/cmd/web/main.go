package main

import (
	"fmt"
	"net/http"

	"github.com/Riyaz-khan-shuvo/go-practice/2-11-app-config/pkg/config"
	"github.com/Riyaz-khan-shuvo/go-practice/2-11-app-config/pkg/handler"
	"github.com/Riyaz-khan-shuvo/go-practice/2-11-app-config/pkg/render"
)

const portNumber string = ":3555"

func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplatePage()
	if err != nil {
		fmt.Println(err)
	}

	app.TemplateCache = tc

	render.NewTemplate(&app)

	fmt.Println("Listing the port http://localhost" + portNumber)
	http.HandleFunc("/", handler.HomePage)
	http.HandleFunc("/about", handler.AboutPage)
	http.ListenAndServe(portNumber, nil)
}
