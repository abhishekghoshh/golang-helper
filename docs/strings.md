# Strings

## String and Runes

A Go string is a read-only slice of bytes storing UTF-8 encoded text. The concept of a "character" is called a **rune** — an int32 representing a Unicode code point.

```go
func StringAndRunes() {
    const s = "hello"

    // len() gives raw byte count
    fmt.Println("Len:", len(s)) // 5

    // Indexing gives raw byte values
    for i := 0; i < len(s); i++ {
        fmt.Printf("%x ", s[i]) // 68 65 6c 6c 6f
    }

    // Count runes (Unicode code points)
    fmt.Println("Rune count:", utf8.RuneCountInString(s))

    // range decodes each rune with its byte offset
    for idx, runeValue := range s {
        fmt.Printf("%#U starts at %d\n", runeValue, idx)
    }

    // Manual decode with utf8.DecodeRuneInString
    for i, w := 0, 0; i < len(s); i += w {
        runeValue, width := utf8.DecodeRuneInString(s[i:])
        fmt.Printf("%#U starts at %d\n", runeValue, i)
        w = width
        examineRune(runeValue)
    }
}
```

## String Builder (Performance)

Concatenating with `+` creates a new string each time (O(n²)). Use `strings.Builder` or `bytes.Buffer` for efficient building.

```go
const STRING_LENGTH = 100000

func StringBuilers() {
    // Normal string (slow)
    calculateTimeTaken("normal string", func() {
        var s string
        for i := 0; i < STRING_LENGTH; i++ {
            s += "x"
        }
    })

    // Byte buffer (fast)
    calculateTimeTaken("byte buffer", func() {
        var buffer bytes.Buffer
        for i := 0; i < STRING_LENGTH; i++ {
            buffer.WriteString("x")
        }
        _ = buffer.String()
    })

    // String builder (fastest)
    calculateTimeTaken("string builder", func() {
        var builder strings.Builder
        for i := 0; i < STRING_LENGTH; i++ {
            builder.WriteString("x")
        }
        _ = builder.String()
    })
}
```

## String Functions

```go
import "strings"

func main() {
    fmt.Println(strings.Contains("test", "es"))      // true
    fmt.Println(strings.Count("test", "t"))           // 2
    fmt.Println(strings.HasPrefix("test", "te"))      // true
    fmt.Println(strings.HasSuffix("test", "st"))      // true
    fmt.Println(strings.Index("test", "e"))           // 1
    fmt.Println(strings.Join([]string{"a","b"}, "-")) // "a-b"
    fmt.Println(strings.Repeat("a", 5))               // "aaaaa"
    fmt.Println(strings.Replace("fooo", "o", "0", -1)) // "f000"
    fmt.Println(strings.Replace("fooo", "o", "0", 1))  // "f0oo"
    fmt.Println(strings.Split("a-b-c", "-"))           // [a b c]
    fmt.Println(strings.ToLower("TEST"))               // "test"
    fmt.Println(strings.ToUpper("test"))               // "TEST"
}
```

## String Formatting

Go uses `fmt.Printf` with formatting verbs.

| Verb | Description | Example Output |
|---|---|---|
| `%v` | Default format | `{1 2}` |
| `%+v` | Struct with field names | `{x:1 y:2}` |
| `%#v` | Go-syntax representation | `main.point{x:1, y:2}` |
| `%T` | Type of the value | `main.point` |
| `%t` | Boolean | `true` |
| `%d` | Integer (base 10) | `123` |
| `%b` | Binary | `1110` |
| `%c` | Character | `!` |
| `%x` | Hex encoding | `1c8` |
| `%f` | Float | `78.900000` |
| `%e` / `%E` | Scientific notation | `1.234000e+08` |
| `%s` | String | `"string"` |
| `%q` | Double-quoted string | `"\"string\""` |
| `%p` | Pointer address | `0xc000010230` |
| `%6d` | Width (right-justified) | `    12` |
| `%-6d` | Width (left-justified) | `12    ` |
| `%6.2f` | Width + precision | `  1.20` |

```go
type point struct{ x, y int }
p := point{1, 2}

fmt.Printf("struct: %v\n", p)
fmt.Printf("struct+: %+v\n", p)         // {x:1 y:2}
fmt.Printf("struct#: %#v\n", p)         // main.point{x:1, y:2}
fmt.Printf("type: %T\n", p)             // main.point
fmt.Printf("bool: %t\n", true)
fmt.Printf("int: %d\n", 123)
fmt.Printf("bin: %b\n", 14)
fmt.Printf("char: %c\n", 33)            // !
fmt.Printf("hex: %x\n", 456)
fmt.Printf("float: %f\n", 78.9)
fmt.Printf("width: |%6d|%6d|\n", 12, 345)
fmt.Printf("width: |%-6.2f|%-6.2f|\n", 1.2, 3.45)

// Sprintf returns a formatted string
s := fmt.Sprintf("sprintf: a %s", "string")

// Fprintf writes to an io.Writer
fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
```

## Regular Expressions

```go
import "regexp"

func RegularExperssions() {
    // Basic match
    match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
    fmt.Println(match) // true

    // Compile for reuse
    r, _ := regexp.Compile("p([a-z]+)ch")

    fmt.Println(r.MatchString("peach"))                    // true
    fmt.Println(r.FindString("peach punch"))               // "peach"
    fmt.Println(r.FindStringIndex("peach punch"))          // [0 5]
    fmt.Println(r.FindStringSubmatch("peach punch"))       // [peach ea]
    fmt.Println(r.FindAllString("peach punch pinch", -1))  // [peach punch pinch]
    fmt.Println(r.FindAllString("peach punch pinch", 2))   // [peach punch]

    // MustCompile panics on error (safe for globals)
    r = regexp.MustCompile("p([a-z]+)ch")

    // Replace
    fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))  // "a <fruit>"

    // Replace with function
    out := r.ReplaceAllFunc([]byte("a peach"), bytes.ToUpper)
    fmt.Println(string(out)) // "a PEACH"
}
```

## Text Templates

```go
import "text/template"

func main() {
    // Create and parse template
    t1 := template.New("t1")
    t1, _ = t1.Parse("Value is {{.}}\n")
    t1.Execute(os.Stdout, "some text") // Value is some text
    t1.Execute(os.Stdout, 5)           // Value is 5

    // Struct fields: {{.FieldName}}
    t2 := template.Must(template.New("t2").Parse("Name: {{.Name}}\n"))
    t2.Execute(os.Stdout, struct{ Name string }{"Jane Doe"}) // Name: Jane Doe

    // Maps
    t2.Execute(os.Stdout, map[string]string{"Name": "Mickey"})

    // Conditionals with whitespace trimming
    t3 := template.Must(template.New("t3").Parse("{{if . -}} yes {{else -}} no {{end}}\n"))
    t3.Execute(os.Stdout, "not empty") // yes
    t3.Execute(os.Stdout, "")          // no

    // Range
    t4 := template.Must(template.New("t4").Parse("Range: {{range .}}{{.}} {{end}}\n"))
    t4.Execute(os.Stdout, []string{"Go", "Rust", "C++"}) // Range: Go Rust C++
}
```
