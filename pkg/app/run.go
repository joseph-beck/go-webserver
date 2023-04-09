package app

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"
)

func Run() {
	welcome := Welcome{
		Name: "Anonymous",
		Time: time.Now().Format(time.Stamp),
	}

	templates := template.Must(template.ParseFiles("templates/index.html"))

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static")))) 
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if name := r.FormValue("name"); name != "" {
			welcome.Name = name
		}
		
		if err := templates.ExecuteTemplate(w, "index.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	log.Println("Serving on port 8080")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
