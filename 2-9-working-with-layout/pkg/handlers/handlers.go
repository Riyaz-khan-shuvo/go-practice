package handlers

import (
	"net/http"

	"github.com/Riyaz-khan-shuvo/go-practice/2-9-working-with-layout/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.gohtml")
}
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.gohtml")
}
