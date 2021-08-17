package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/hngi8/task4/controller"
)

func main() {

	PORT := os.Getenv("PORT")
	if PORT == ""  {
		PORT = "3000"
	}

	http.HandleFunc("/", controller.Contact)
	fs := http.FileServer(http.Dir("view/index.html"))
	http.Handle("/contact", http.StripPrefix("/view/index.html", fs))



	fmt.Println("Server Started on port"+PORT+"....")
	http.ListenAndServe(":"+PORT, nil)
}
