package controller

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hngi8/task4/models"
)

var tmpl = template.Must(template.ParseFiles("view/index.html"))

func Contact(w http.ResponseWriter, r *http.Request) {

	// Connect database
	db, err := sql.Open("mysql", "b7ca0ef21f6a3e:4c49d0c9@tcp(s-cdbr-east-04.cleardb.com)/heroku_da5272d9c3c8ace?")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("Database connected successfullly")

	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}
	data := models.Contact{
		Name:    r.FormValue("name"),
		Email:   r.FormValue("email"),
		Subject: r.FormValue("subject"),
		Message: r.FormValue("message"),
	}
	_ = data
	//insert message database contact
	newContact := new(models.Contact)
	newContact.Name = r.FormValue("name")
	newContact.Email = r.FormValue("email")
	newContact.Subject = r.FormValue("subject")
	newContact.Message = r.FormValue("message")

	insert, err := db.Query("INSERT INTO users (firstname, lastname, email) VALUES ('abiola', 'fasanya', 'abiola@mail.com')")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println("Data Inserted Successfullly")
	//mysql: //b7ca0ef21f6a3e:4c49d0c9@us-cdbr-east-04.cleardb.com/heroku_da5272d9c3c8ace?
	tmpl.Execute(w, struct{ Success bool }{true})
}
