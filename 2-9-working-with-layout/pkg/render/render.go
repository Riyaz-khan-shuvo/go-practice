package render

import (
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"
)

var functions = template.FuncMap{}

func RenderTemplate(w http.ResponseWriter, tem string) {
	pToTem, err := template.ParseFiles("../../template/" + tem)

	if err != nil {
		fmt.Println(err)
	}
	pToTem.Execute(w, nil)

	RenderTemplateTest(w)
}

func RenderTemplateTest(w http.ResponseWriter) (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("../../template/*.gohtml")

	if err != nil {
		fmt.Println(err)
	}
	for _, page := range pages {
		name := filepath.Base(page)

		// template set

		fmt.Println("Currently page is : ", page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			fmt.Println(err)
		}

		matches, err := filepath.Glob("../../template/.*layouthtml")
		if err != nil {
			fmt.Println(err)
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("../../template/.*layouthtml")
			if err != nil {
				fmt.Println(err)
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil

}
