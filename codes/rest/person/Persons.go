package person

type Person struct {
	Id        int    `json:"id" xml:"id"`
	FirstName string `json:"firstName" xml:"firstName"`
	LastName  string `json:"lastName" xml:"lastName"`
	Age       int    `json:"age" xml:"age"`
	Gender    string `json:"gender" xml:"gender"`
}
