package dto

type Person struct {
	ID       uint   `json:"id"`
	Name     string `json:"name,omitempty"`
	UserName string `json:"userName,omitempty"`
	Age      int64  `json:"age,omitempty"`
	Email    string `json:"email,omitempty"`
	Notes    []Note `json:"notes"`
}

type Note struct {
	NoteId  uint   `json:"noteId"`
	Content string `json:"content,omitempty"`
}
