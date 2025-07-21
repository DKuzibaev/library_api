package models

type Book struct {
	ID          string `json:"id,omitempty"` // omitempty - поле не будет включено в JSON, если оно пустое
	Author      string `json:"author"`
	Label       string `json:"label"`
	Description string `json:"description"`
	Genre       string `json:"genre"`
	WrittenAt   string `json:"written_at"`
	Count       uint   `json:"count,omitempty"` // Количество книг
}
