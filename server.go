package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/hngi8/controller"
)

func main() {

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "3000"
	}

	http.HandleFunc("/", controller.Index)

	http.HandleFunc("/create-table", controller.ContactTable)

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	fmt.Println("Server Started on port" + PORT + "....")
	http.ListenAndServe(":"+PORT, nil)
}
