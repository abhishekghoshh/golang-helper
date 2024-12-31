package mongodb

type Person struct {
	Id     int64
	Name   string
	Age    int64
	Email  string
	UserId string
	Notes  map[string]any
}
