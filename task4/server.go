package main

import (
	"fmt"
	"net/http"

	"example.com/hngi8/task4/controller"
)

func main() {

	http.HandleFunc("/contact", controller.Contact)

	// fs := http.FileServer(http.Dir("view/index.html"))
	// http.Handle("/home", http.StripPrefix("/view/index.html", fs))

	fmt.Println("Server Started....")
	http.ListenAndServe(":3000", nil)
}
