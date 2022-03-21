package render

import (
	"bytes"
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/Riyaz-khan-shuvo/go-practice/2-13-optimizing-the-template/pkg/config"
)

var functions = template.FuncMap{}

var app *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	app = a
}

func RenderPage(w http.ResponseWriter, temp string) {

	tc := app.TemplateCache
	// tc, err := CreateTemplate()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	t, ok := tc[temp]

	buf := new(bytes.Buffer)
	_ = t.Execute(buf, nil)
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println(err)
	}

	if !ok {
		fmt.Println(ok)
	}

	// pToTem, err := template.ParseFiles("../../template/" + temp)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// pToTem.Execute(w, nil)
}

func CreateTemplate() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("../../template/*.gohtml")
	if err != nil {
		fmt.Println(err)
	}
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			fmt.Println(err)
		}
		matches, err := filepath.Glob("../../template/*.temhtml")
		if err != nil {
			fmt.Println(err)
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("../../template/*.temhtml")
			if err != nil {
				fmt.Println(err)
			}
		}
		myCache[name] = ts
	}

	return myCache, nil
}
