package msg

import "fmt"

func Greet(n string, f func(string)) {
	f(n)
}

func Welcome(n string) {
	fmt.Printf("%s, You are welcome to Hngi8 Internship, and this is your first task", n)
}