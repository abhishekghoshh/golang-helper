package iterator

import (
	"design-pattern/src/behavioral/iterator/list"
)

func Main() {
	user1 := &User{
		name: "Abhishek",
		age:  25,
	}
	user2 := &User{
		name: "Nasim",
		age:  26,
	}
	user3 := &User{
		name: "Bishal",
		age:  26,
	}

	userList := list.New(user1, user2, user3)
	Print(&userList)
	userList.Add(&User{name: "Niket", age: 25})
	Print(&userList)
}

func Print(collection list.Collection[User]) {
	iterator := collection.Iterator()
	for iterator.HasNext() {
		user := iterator.Next()
		user.Print()
	}
}
