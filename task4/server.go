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
	http.FileServer(http.Dir("view"))
	http.Handle("/", http.FileServer(http.Dir("view")))

	http.HandleFunc("/", controller.Contact)


	fmt.Println("Server Started on port"+PORT+"....")
	http.ListenAndServe(":"+PORT, nil)
}
