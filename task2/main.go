package main

import (
	"fmt"

	"github.com/hngi8/task2/msg"
)

func main() {
	// Print out your name
	fmt.Println("Name: Abiola")
	msg.Greet("Abiola", msg.Welcome)

	// for web

	// func HandleFunc(w http.ResponseWriter, r *http.Request)
}
