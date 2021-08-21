package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hngi8/config"
	"github.com/hngi8/models"
	_ "github.com/lib/pq"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("*.html"))
}

func Index(w http.ResponseWriter, r *http.Request) {
	profile := models.Profile{
		Name:     "Abiola Fasanya",
		Email:    "harbiola78@gmail.com",
		Phone:    "2348102307473",
		Location: "Lagos, Nigeria",
	}

	req := models.Contact{
		Name:    r.FormValue("name"),
		Email:   r.FormValue("email"),
		Subject: r.FormValue("subject"),
		Message: r.FormValue("message"),
	}

	out := struct {
		Success bool
		Profile models.Profile
		Message map[string]string
	}{false, profile, nil}

	if r.Method == "GET" {
		t, _ := template.ParseFiles("index.html")
		t.Execute(w, out)
		return
	}

	info := models.Info{Message: req.Name + " Thank You for Your Message it has been received successfully!"}

	msg := struct {
		Success bool
		Profile models.Profile
		Message models.Info
	}{
		Success: true,
		Profile: profile,
		Message: info,
	}

	

	env := os.Getenv("ENVIRONMENT")
	if env == "production" {
		addContact(req.Name, req.Email, req.Subject, req.Message)
	}else{
		// insert message database contact local environment
		db := config.DbConn()
		insert := "INSERT INTO messages (`name`, `email`, `subject`, `message`) VALUES (?,?,?,?)"

		q, err := db.Query(insert, req.Name, req.Email, req.Subject, req.Message)
		if err != nil {
			log.Fatal(err, "An error occured!")
		}
		fmt.Println("Data Inserted Successfullly")
		defer q.Close()
	}

	fmt.Println(msg.Profile, msg.Success)
	tmpl.Execute(w, msg)

	// http.Redirect(w, r, "/", http.StatusSeeOther)
}

func ContactTable(w http.ResponseWriter, r *http.Request) {
	db := config.DbConn()
	table := `
		CREATE TABLE IF NOT EXISTS contact (
			id serial PRIMARY KEY,
			name varchar(111) NOT NULL,
			email varchar(111) NOT NULL,
			subject varchar(111) NOT NULL,
			message varchar(255) NOT NULL, 
                        created_at timestamp with time zone DEFAULT current_timestamp
                 )`

	d, err := db.Query(table)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err.Error())
	}
	_ = d
	fmt.Println("Table created successfully")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func addContact(a, b, c, d string) {
	db := config.DbConn()
	sqlStatement := `
		INSERT INTO contact (name, email, subject, message)
		VALUES ($1, $2, $3, $4)`

	_, err := db.Exec(sqlStatement, a, b, c, d)
	if err != nil {
		panic(err)
	}
	fmt.Println("Data Inserted Successfullly")
}
