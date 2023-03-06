package main

import "fmt"

func main() {
	mp := make(map[string]int)
	mp["abhishek"] = 25
	mp["kushal"] = 25
	mp["nasim"] = 26
	fmt.Println(mp["abhishek"])
	fmt.Println(mp["pal"])

	m := map[string]int{
		"James": 42,
		"Amy":   24}

	fmt.Println(m["Amy"])
	delete(m, "James")
	fmt.Println(m)
}
