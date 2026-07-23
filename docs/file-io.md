# Reading and Writing Files

## Reading Files

```go
func ReadingFiles() {
    // Slurp entire file into memory
    dat, err := os.ReadFile("./data/data.txt")
    if err != nil { panic(err) }
    fmt.Print(string(dat))

    // Open file for controlled reading
    f, err := os.Open("./data/data.txt")
    if err != nil { panic(err) }
    defer f.Close()

    // Read specific number of bytes
    b1 := make([]byte, 5)
    n1, err := f.Read(b1)
    fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

    // Seek to a position and read
    f.Seek(6, 0) // offset 6, relative to start (0=start, 1=current, 2=end)
    b2 := make([]byte, 2)
    n2, _ := f.Read(b2)
    fmt.Printf("%d bytes @ 6: %s\n", n2, string(b2[:n2]))

    // Read with io.ReadAtLeast (ensures minimum bytes)
    f.Seek(6, 0)
    b3 := make([]byte, 2)
    io.ReadAtLeast(f, b3, 2)

    // Rewind: Seek(0, 0)
    f.Seek(0, 0)

    // Buffered reader
    r4 := bufio.NewReader(f)
    b4, _ := r4.Peek(5)
    fmt.Printf("5 bytes (peek): %s\n", string(b4))

    f.Close()
}
```

## Print File (cat)

```go
func CatFile() {
    file, err := os.Open(filename)
    if err != nil { panic(err) }
    defer file.Close()
    io.Copy(os.Stdout, file)
}
```

## Writing Files

```go
func WritingFiles() {
    // Simple: dump string/bytes to file
    d1 := []byte("hello\ngo\n")
    os.WriteFile("./data/newData.txt", d1, 0644)

    // Granular writes with os.Create
    f, _ := os.Create("./data/newData-2.txt")
    defer f.Close()

    f.Write([]byte{115, 111, 109, 101, 10}) // "some\n"
    f.WriteString("writes\n")
    f.Sync() // flush to stable storage

    // Buffered writer
    w := bufio.NewWriter(f)
    w.WriteString("buffered\n")
    w.Flush()
}
```

## Appending to Files

```go
func appendToFile(filename, newData string) string {
    f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
    if err != nil { panic(err) }
    defer f.Close()

    f.WriteString(newData + "\n")

    data, _ := os.ReadFile(filename)
    return string(data)
}
```

## Line Filters (stdin → stdout)

```go
func LineFilters() {
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        ucl := strings.ToUpper(scanner.Text())
        fmt.Println(ucl)
    }
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "error:", err)
        os.Exit(1)
    }
}
// Usage: cat data.txt | go run main.go
```

## File Paths

```go
func FilePaths() {
    // Portable path joining
    path := filepath.Join("dir1", "dir2", "filename")

    // Normalizes separators and ..
    filepath.Join("dir1//", "filename")       // "dir1/filename"
    filepath.Join("dir1/../dir1", "filename") // "dir1/filename"

    // Split path
    filepath.Dir(path)                    // "dir1/dir2"
    filepath.Base(path)                   // "filename"
    dir, base := filepath.Split(path)     // "dir1/dir2/", "filename"

    // Absolute path check
    filepath.IsAbs("/dir/file")           // true
    filepath.IsAbs("dir/file")            // false

    // File extension
    ext := filepath.Ext("config.json")    // ".json"

    // Relative path between two paths
    rel, _ := filepath.Rel("a/b", "a/b/t/file")  // "t/file"
}
```

## Directories

```go
func Directories() {
    // Create directory
    os.Mkdir("subdir", 0755)
    defer os.RemoveAll("subdir")

    // Create directory hierarchy (like mkdir -p)
    os.MkdirAll("subdir/parent/child", 0755)

    // List directory
    entries, _ := os.ReadDir("subdir/parent")
    for _, e := range entries {
        fmt.Println(e.Name(), e.IsDir())
    }

    // Change directory
    os.Chdir("subdir/parent/child")
    os.Chdir("../../..")

    // Walk directory recursively
    filepath.Walk("subdir", func(path string, info os.FileInfo, err error) error {
        fmt.Println(" ", path, info.IsDir())
        return nil
    })
}
```

## Temporary Files and Directories

```go
func TemporaryFilesAndDirectories() {
    // Create temp file (in OS default temp dir)
    f, _ := os.CreateTemp("", "sample")
    fmt.Println("Temp file:", f.Name()) // /tmp/sample12345
    defer os.Remove(f.Name())
    f.Write([]byte{1, 2, 3, 4})

    // Create temp directory
    dname, _ := os.MkdirTemp("", "sampledir")
    defer os.RemoveAll(dname)

    // Write file inside temp dir
    fname := filepath.Join(dname, "file1")
    os.WriteFile(fname, []byte{1, 2}, 0666)
}
```

## Embed Directive

The `//go:embed` compiler directive embeds files/folders into the Go binary at build time.

```go
import _ "embed"

// Embed single file as string
//go:embed embed/single_file.txt
var fileString string

// Embed as []byte
//go:embed embed/single_file.txt
var fileByte []byte

// Embed entire folder as virtual filesystem
//go:embed embed/*
var folder embed.FS

func EmbeddedDirective() {
    fmt.Println(fileString)
    println(string(fileByte))

    content1, _ := folder.ReadFile("embed/file1.hash")
    println(string(content1))
}
```

Setup:
```bash
mkdir -p files/embed
echo "random file content" > files/embed/single_file.txt
echo "123" > files/embed/file1.hash
echo "456" > files/embed/file2.hash
```
