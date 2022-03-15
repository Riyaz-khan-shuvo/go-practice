package handler

import (
	"net/http"

	"github.com/Riyaz-khan-shuvo/go-practice/2-11-app-config/pkg/render"
)

func HomePage(w http.ResponseWriter, r *http.Request) {

	render.RenderPages(w, "home.gohtml")

}

func AboutPage(w http.ResponseWriter, r *http.Request) {

	render.RenderPages(w, "about.gohtml")

}
