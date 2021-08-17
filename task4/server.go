package main

import (
	"fmt"
	"net/http"

	"github.com/fastbeetech/hngi8/task4/controller"
)

func main() {

	http.FileServer(http.Dir("public/css"))
	// http.Handle("/css", http.StripPrefix("/public/css/", fs))

	http.HandleFunc("/", controller.Contact)


	fmt.Println("Server Started....")
	http.ListenAndServe(":3000", nil)
}
