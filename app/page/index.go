package page

import (
	"encoding/json"
	"fmt"
	"github.com/idasilva/project-web/external"
	"github.com/idasilva/project-web/validator"
	"html/template"
	"log"
	"net/http"
	"time"
)

var temp = template.Must(template.ParseGlob("./etc/tampletes/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	c := external.NewClient()

	resp, err := c.Request(nil)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	var documents = []external.RequestNews{}
	err = json.NewDecoder(resp.Body).Decode(&documents)
	if err != nil {
		log.Fatal(err)
	}
	for i, d := range documents{
		dia , mes, ano:= d.CreatedAt.Date()
		m :=  mes.String()
		d.Data =  fmt.Sprintf("%v-%s-%v",ano,m, dia)
		documents[i].Data = d.Data
	}
	temp.ExecuteTemplate(w, "Index", documents)
}
func Form(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Formulario", nil)
}

func Register(w http.ResponseWriter, r *http.Request) {

	doc := &external.RequestNews{}
	doc.Title = r.FormValue("title")
	doc.Body = r.FormValue("body")
	doc.CreatedAt = time.Now()
	doc.UpdatedAt = time.Now()

	valid := validator.NewValidate()

	Isvalid, err := valid.ValidateStruct(doc)
	if err != nil || !Isvalid {
		http.Redirect(w, r, "/form/1", 301)
		return
	}
	c := external.NewClient()

	resp, err := c.Request(doc)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK{
		http.Redirect(w, r, "/form/1", 301)
		return

	}

	http.Redirect(w, r, "/index", 301)
}
