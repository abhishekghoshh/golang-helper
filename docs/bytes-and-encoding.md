# Bits, Bytes, and Byte Slices

## Bits and Bytes

A bit references a single binary value, either a 1 or a 0. For a computer to process information the information must be written in binary. If all we could send to a computer would be a single 1 or a 0, we wouldn't be as far as we are today with computers. Luckily, bits can be added together to reference much larger numbers than 1 and 0.

Bits can be added together, and their values are equal to 2 to the power of the place of the bit (from right to left). When the bit is 1, the bit value is added to the binary value. When the bit is 0, the value is not added. The binary value consisting of 8 bits equal to `00000001` would equal 1 (2 to the 0th power = 1), the value `00000010` would equal 2 (2 to the 1st power), and the value `00000100` would equal 4 (2 to the 2nd power). Once you start combining these, you start to see how they can soon represent a large number. The value of `00000011` is 3 (2 to the 0th power is 1 + 2 to the 1st power is 2 = 3), the value of `10010010` is 146 (2 to the 2nd power + 2 to the 4th power + 2 to the 7th power) and `11111111` is 256.

## Unicode and UTF-8

Before diving into byte slices, we need to briefly visit encoding. Unicode is "the universal character encoding". Unicode supports over 137,000 characters and 150 different languages. Each character has a specific Unicode code point that represents the character. For example, the Unicode code point for the capital letter "T" is U+0054, the Unicode code point for the lowercase letter "t" is U+0074, and the Unicode point for the lower left triangle "◺" is U+25FA.

The "problem" with Unicode is that it is most often represented in software as an int32 (32-bit integer). Most characters in widespread use could fit into a much smaller number than required for a 32-bit data type. This is where UTF-8 enters the picture.

UTF-8 is a variable length encoding of Unicode code points. For each Unicode code point, it uses between 1 and 4 bytes. All the most common characters can be represented using 1-2 bytes (all ASCII characters can be represented in 1 byte). UTF-8 allows us to use all the characters defined by Unicode but allows us to save some extra space and only reach for the 3rd or 4th byte when we really need it. To again list the same examples, we could represent the Unicode code point "T" as 84, "t" as 116, and "◺" as 226 151 186.

This is a very brief overview of encoding, and I would highly recommend reading the article:
[The Absolute Minimum Every Software Developer Absolutely, Positively Must Know About Unicode and Character Sets](https://www.joelonsoftware.com/2003/10/08/the-absolute-minimum-every-software-developer-absolutely-positively-must-know-about-unicode-and-character-sets-no-excuses/).

## Byte Slices

Byte slices are a list of bytes that represent UTF-8 encodings of Unicode code points. Taking the information from above, we could create a byte slice that represents the word "Go":

```go
bs := []byte{71, 111}
fmt.Printf("%s", bs) // Output: Go
```

You may notice the `%s` used here. This converts the byte slice to a string. Strings are literally made up of **arrays of bytes**. This makes converting strings to byte slices and byte slices to strings very easy and convenient. `%d` is also often used with byte slices and prints the UTF-8 decimal value of each byte.

```go
s := "Wow look at me"
bs := []byte(s)
fmt.Printf("%s", bs) // Output: Wow look at me
fmt.Printf("%d", bs) // Output: [87 111 119 32 108 111 111 107 32 97 116 32 109 101]
```

To view something a little more interesting, here is an example of a byte slice representing a non-ASCII value:

```go
bs := []byte("◺")
fmt.Println(bs) // Output: [226 151 186]
s := string(bs)
fmt.Println(len(s)) // Output: 3
```

The length of this string may seem confusing at first, but remember a string is made up of a byte slice array. The length of this string is 3 because the UTF-8 value of "◺" is 226 151 186, which means that UTF-8 uses 3 bytes to represent the Unicode code point. To get the Unicode code point (or rune) count in a string, we can use the utf8 package in the standard library:

```go
import (
    "fmt"
    "unicode/utf8"
)

bs := []byte("◺") 
s := string(bs)

fmt.Println(utf8.RuneCountInString(s)) // Output: 1
```

This is important to remember if Unicode character count is important to your software. Luckily, we do not need to handle this within range, because Go will implicitly loop over strings by their Unicode code points. The only catch here is that the index will still be incremented by the string's byte count. Notice how `i` in the below example jumps from 3 to 6 after printing the triangle:

```go
func main() {
 for i, b := range "Hi ◺ there" {
  fmt.Printf("i: %d. b: %q\n", i, b)
 }
}
// i: 0. b: 'H'
// i: 1. b: 'i'
// i: 2. b: ' '
// i: 3. b: '◺'
// i: 6. b: ' '
// i: 7. b: 't'
// i: 8. b: 'h'
// i: 9. b: 'e'
// i: 10. b: 'r'
// i: 11. b: 'e'
```

Using `%c` here prints the character represented by the corresponding Unicode code point. Since each of the characters in the strings of the rune type (Unicode code point) we cannot use `%s` to print here.

Going back to the example I gave earlier: `string(42) // Output: *`. This may start to be making some sense. The UTF-8 decimal value for `*` is 42. So when we pass the integer 42 into a string, it is creating a byte slice containing 42, which is equal to `*`.

## Byte Slices vs Strings

So now that we know about byte slices, when should we use them over strings? The main difference in strings and byte slices in Go is the way they are handled in memory. **Strings are immutable**, meaning they cannot be changed within memory. So, every time you add or remove from a string, Go is creating a brand-new string in memory. On the contrary, **byte slices are mutable**. So, whenever you are adding to a byte slice, you are not creating a new object in memory. Depending on the circumstance, this can affect application speed.

## Additional Resources

- [The Go Programming Language by Alan A. A. Donovan and Brian W. Kernighan (specifically chapter 2 section 3.5)](https://www.amazon.com/Programming-Language-Addison-Wesley-Professional-Computing/dp/0134190440/)
- [The Absolute Minimum Every Software Developer Absolutely, Positively Must Know About Unicode and Character Sets (No Excuses!)](https://www.joelonsoftware.com/2003/10/08/the-absolute-minimum-every-software-developer-absolutely-positively-must-know-about-unicode-and-character-sets-no-excuses/)
- [Unicode Website](https://home.unicode.org/)
- [UTF-8 Character Set](https://www.fileformat.info/info/charset/UTF-8/list.htm)
