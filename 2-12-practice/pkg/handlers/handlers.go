package handlers

import (
	"net/http"

	"github.com/Riyaz-khan-shuvo/go-practice/2-12-practice/pkg/render"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	render.RenderPages(w, "home.gohtml")
}

func AboutPage(w http.ResponseWriter, r *http.Request) {
	render.RenderPages(w, "about.gohtml")
}
