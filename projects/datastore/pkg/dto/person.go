package dto

type Person struct {
	ID       string `json:"id"`
	Name     string `json:"name,omitempty"`
	UserName string `json:"userName,omitempty"`
	Age      int64  `json:"age,omitempty"`
	Email    string `json:"email,omitempty"`
	Notes    []Note `json:"notes"`
}

type Note struct {
	NoteId  string `json:"noteId"`
	Content string `json:"content,omitempty"`
}
