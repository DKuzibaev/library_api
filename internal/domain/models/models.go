package models

type Book struct {
	ID          string `json:"id"`
	Author      string `json:"author"`
	Label       string `json:"label"`
	Description string `json:"description"`
	Genre       string `json:"genre"`
	WrittenAt   int    `json:"written_at"`
	Count       uint   `json:"count"`
}
