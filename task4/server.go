package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/hngi8/task4/controller"
)

func main() {

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "3000"
	}

	http.HandleFunc("/contact", controller.Index)

	http.HandleFunc("/", controller.Contact)

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	fmt.Println("Server Started on port" + PORT + "....")
	http.ListenAndServe(":"+PORT, nil)
}
