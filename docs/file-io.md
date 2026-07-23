# Reading and Writing Files

## Reading Files

In order to read from files on your local filesystem, you'll have to use the `io/ioutil` module. You'll first have to pull the contents of a file into memory by calling `ioutil.ReadFile("/path/to/file.ext")` which will take in the path to the file you wish to read in as its only parameter. This will return either the data of the file, or an err which can be handled as you normally handle errors in go.

```go
func retrieveFromFile(filename string) string {
    bs, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Println("unable to get the file", err)
        return ""
    }
    return string(bs)
}
```

## Writing Files

In order to write content to files using Go, we'll again have to leverage the io/ioutil module. We'll first have to construct a byte array that represents the content we wish to store within our files.

```go
mydata := []byte("all my data I want to write to a file")
```

Once we have constructed this byte array, we can then call `ioutil.WriteFile()` to write this byte array to a file. The `WriteFile()` method takes in 3 different parameters: the first is the location of the file we wish to write to, the second is our mydata object, and the third is the FileMode, which represents our file's mode and permission bits.

```go
func saveToFile(filename, data string) error {
    e := ioutil.WriteFile(filename, []byte(data), 0666)
    if e == nil {
        fmt.Println("data saved to file")
    } else {
        fmt.Println("data is not saved to file")
    }
    return e
}
```

## Appending to Existing Files

```go
func appendToFile(filename, newData string) {
    f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
    if err != nil {
        panic(err)
    }
    defer f.Close()
    if _, err = f.WriteString(newData + "\n"); err != nil {
        panic(err)
    }
}
```
