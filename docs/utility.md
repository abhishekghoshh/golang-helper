# Utility Functions

## Random Numbers

```go
import (
    cryptorand "crypto/rand"
    rand "math/rand/v2"
)

func RandomNumberGenerator() {
    // math/rand/v2: pseudo-random
    fmt.Println(rand.IntN(100))      // 0 <= n < 100
    fmt.Println(rand.Float64() * 5 + 5) // 5.0 <= f < 10.0

    // With known seed (reproducible)
    s := rand.NewPCG(42, 1024)
    r := rand.New(s)
    fmt.Println(r.IntN(100))

    // crypto/rand: cryptographically secure (slower)
    RandomCrypto, _ := cryptorand.Prime(cryptorand.Reader, 10)
    fmt.Println(RandomCrypto)
}
```

## Number Parsing

```go
import "strconv"

func NumberParsing() {
    f, _ := strconv.ParseFloat("1.234", 64)     // 1.234
    i, _ := strconv.ParseInt("123", 0, 64)      // 123 (base inferred)
    d, _ := strconv.ParseInt("0x1c8", 0, 64)    // 456 (hex)
    u, _ := strconv.ParseUint("789", 0, 64)     // 789
    k, _ := strconv.Atoi("135")                 // 135 (convenience)
    _, e := strconv.Atoi("wat")                 // error: invalid syntax
    fmt.Println(e)
}
```

## URL Parsing

```go
import "net/url"

func URLParsing() {
    s := "postgres://user:pass@host.com:5432/path?k=v#f"
    u, _ := url.Parse(s)

    fmt.Println(u.Scheme)              // "postgres"
    fmt.Println(u.User.Username())     // "user"
    p, _ := u.User.Password()
    fmt.Println(p)                     // "pass"

    fmt.Println(u.Host)                // "host.com:5432"
    host, port, _ := net.SplitHostPort(u.Host)
    fmt.Println(host)                  // "host.com"
    fmt.Println(port)                  // "5432"

    fmt.Println(u.Path)                // "/path"
    fmt.Println(u.Fragment)            // "f"

    // Query params
    fmt.Println(u.RawQuery)            // "k=v"
    m, _ := url.ParseQuery(u.RawQuery)
    fmt.Println(m["k"][0])             // "v"
}
```

## SHA256 Hashing

```go
import "crypto/sha256"

func SHA256Hashing() {
    str := "sha256 this string"
    hsh := sha256.New()
    hsh.Write([]byte(str))
    bs := hsh.Sum(nil) // nil means append to nothing
    fmt.Printf("%x\n", bs) // hex output
}
```

## Base64 Encoding

```go
import b64 "encoding/base64"

func Base64Encoding() {
    data := "abc123!?$*&()'-=@~"

    // Standard encoding
    sEnc := b64.StdEncoding.EncodeToString([]byte(data))
    fmt.Println(sEnc)
    sDec, _ := b64.StdEncoding.DecodeString(sEnc)
    fmt.Println(string(sDec))

    // URL-compatible encoding (uses - and _ instead of + and /)
    uEnc := b64.URLEncoding.EncodeToString([]byte(data))
    fmt.Println(uEnc)
    uDec, _ := b64.URLEncoding.DecodeString(uEnc)
    fmt.Println(string(uDec))
}
```

## Reflection

```go
import "reflect"

func Reflections() {
    x := 100
    v := reflect.ValueOf(x)
    t := v.Type()
    fmt.Println("Type:", t) // "Type: int"

    // Get value as interface
    fmt.Println(v.Interface())

    // Check kind
    fmt.Println(v.Kind()) // reflect.Int

    // For structs: iterate fields, get tags
    val := reflect.Indirect(reflect.ValueOf(someStruct))
    for i := 0; i < val.NumField(); i++ {
        typeField := val.Type().Field(i)
        tag := typeField.Tag.Get("validate")
        valueField := val.Field(i)
    }
}
```
