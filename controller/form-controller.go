package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hngi8/config"
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
	_, _, _ = data, profile, tmpl

	db := config.DbConn()

	// insert message database contact
	req := models.Contact{
		Name: r.FormValue("name"),
		Email: r.FormValue("email"),
		Subject: r.FormValue("subject"),
		Message: r.FormValue("message"),
	}

	insert := "INSERT INTO messages (`name`, `email`, `subject`, `message`) VALUES (?,?,?,?)"

	q, err := db.Query(insert, req.Name, req.Email, req.Subject, req.Message)
	if err != nil {
		panic(err.Error())
	}
	log.Printf("Insert Name: %s | Email: %s | Subject: %s | Message: %s", req.Name, req.Email, req.Subject, req.Message)

	defer q.Close() 


	msg := struct {
		Success bool
		Profile models.Profile
	}{
		Success: true,
		Profile: profile,
	}

	fmt.Println(msg.Profile, msg.Success)
	fmt.Println("Data Inserted Successfullly")
	// tmpl.ExecuteTemplate(w, "index.html", msg)
	// tmpl.Execute(w, struct{ Success bool }{true})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
