package controller

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/hngi8/models"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("*.html"))
}

// var tmpl = template.Must(template.ParseFiles("view/index.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	profile := models.Profile{
		Name:     "Abiola Fasanya",
		Email:    "harbiola78@gmail.com",
		Phone:    "2348102307473",
		Location: "Lagos, Nigeria",
	}

	if r.Method != http.MethodPost {
		t, _ := template.ParseFiles("index.html")
		t.Execute(w, profile)
		return
	}
	data := models.Contact{
		Name:    r.FormValue("name"),
		Email:   r.FormValue("email"),
		Subject: r.FormValue("subject"),
		Message: r.FormValue("message"),
	}
	_, _ = data, profile

	msg := struct {
		Success bool
		Profile models.Profile
	}{
		Success: true,
		Profile: profile,
	}

	fmt.Println(msg.Profile, msg.Success)
	fmt.Println("Data Inserted Successfullly")
	tmpl.ExecuteTemplate(w, "index.html", struct{ Success bool }{true})
	// tmpl.Execute(w, struct{ Success bool }{true})
}
