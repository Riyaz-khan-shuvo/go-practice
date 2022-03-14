package render

import (
	"fmt"
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, tem string) {
	pToTem, err := template.ParseFiles("../../template/" + tem)

	if err != nil {
		fmt.Println(err)
	}
	pToTem.Execute(w, nil)

}
