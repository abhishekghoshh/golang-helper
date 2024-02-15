package iterator

import "fmt"

type User struct {
	name string
	age  int
}

func (u User) Print() {
	fmt.Println("name is", u.name, "age is", u.age)
}
