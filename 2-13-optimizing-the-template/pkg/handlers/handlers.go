package handlers

import (
	"net/http"

	"github.com/Riyaz-khan-shuvo/go-practice/2-13-optimizing-the-template/pkg/render"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	render.RenderPage(w, "home.gohtml")
}

func AboutPage(w http.ResponseWriter, r *http.Request) {
	render.RenderPage(w, "about.gohtml")
}
