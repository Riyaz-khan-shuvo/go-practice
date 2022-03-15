package handlers

import (
	"net/http"

	"github.com/Riyaz-khan-shuvo/go-practice/2-10-building-template-cache/pkg/render"
)

func HomePage(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "home.gohtml")
}

func AboutPage(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.gohtml")
}

func ProjectPage(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "project.gohtml")
}
