package main

import (
	"fmt"
	"net/http"
	"text/template"
	"time"
)

type Welcome struct {
	Sale string
	Time string
}

func main() {
	Welcome := Welcome{"Sale Begins now'", time.Now().Format(time.Stamp)}
	template := template.Must(template.ParseFiles("template/template.html"))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if sale := r.FormValue("sale"); sale != "" {
			Welcome.Sale = sale
		}

		if err := template.ExecuteTemplate(w, "template.html", Welcome); err != nil {
			// http.Error automatically sets the appropriate status code
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	fmt.Println(http.ListenAndServe(":8000", nil))

}
