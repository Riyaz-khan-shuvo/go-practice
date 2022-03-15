package render

import (
	"bytes"
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"
)

var functions = template.FuncMap{}

func RenderTemplate(w http.ResponseWriter, temp string) {

	tc, err := CreateTemplate()
	if err != nil {
		fmt.Println(err)
	}
	t, ok := tc[temp]
	if !ok {
		fmt.Println(ok)
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, nil)
	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println(err)
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
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("../../template/*.temphtml")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts.ParseGlob("../../template/*.temphtml")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}

	return myCache, nil
}
