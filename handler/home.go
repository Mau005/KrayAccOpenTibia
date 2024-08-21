package handler

import (
	"log"
	"net/http"
	"text/template"
)

type HomeHandler struct{}

func (hh *HomeHandler) GetHome(w http.ResponseWriter, r *http.Request) {
	templ, err := template.New("index.html").ParseFiles("www/index.html")
	if err != nil {
		log.Println("error create template", err)
		return
	}
	err = templ.Execute(w, nil)
	if err != nil {
		log.Println("error execute template", err)
		return
	}
}
