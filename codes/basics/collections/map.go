package collections

import "fmt"

func Maps() {
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

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
	printMap(colors)
}
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
