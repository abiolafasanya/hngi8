package models

type Contact struct {
	Name    string
	Subject string
	Email   string
	Message string
}

type Info struct {
    Message string `json:"message"`
}
