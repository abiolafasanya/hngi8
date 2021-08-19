package models

type Profile struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Location string `json:"location"`
}

type Project struct {
	Name string `json:"name"`
	Url string `json:"url"`
	Github string `json:"github"`
}