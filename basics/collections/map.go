package collections

import "fmt"

/*
Maps are Go's built-in associative data type (sometimes called hashes or dicts in other languages)
https://en.wikipedia.org/wiki/Associative_array

To create an empty map, use the builtin make: make(map[key-type]val-type)
Set key/value pairs using typical name[key] = val syntax
Get a value for a key with name[key]
If the key doesn't exist, the zero value of the value type is returned.

The builtin len returns the number of key/value pairs when called on a map.
*/

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

	fmt.Println("current map is", m)

	// The builtin delete removes key/value pairs from a map.
	delete(m, "James")
	fmt.Println("after deleting the James key", m)

	// To remove all key/value pairs from a map, use the clear builtin
	clear(m)
	fmt.Println("after clearing the entire map:", m)

	// The optional second return value when getting a value from a map
	// indicates if the key was present in the map. This can be used to disambiguate
	// between missing keys and keys with zero values like 0 or "". Here we didn't
	// need the value itself, so we ignored it with the blank identifier _.
	_, isKeyPresent := m["k2"]
	fmt.Println("isKeyPresent:", isKeyPresent)

	// You can also declare and initialize a new map in the same line with this syntax.
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
