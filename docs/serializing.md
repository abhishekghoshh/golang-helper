# Serializing (JSON & XML)

## JSON

Go's `encoding/json` package handles encoding and decoding. Only exported fields (uppercase) are included. Use struct tags to customize field names.

```go
type response1 struct {
    Page   int
    Fruits []string
}

type response2 struct {
    Page   int      `json:"page"`
    Fruits []string `json:"fruits"`
}

func TestJSONHandling() {
    // Marshal basic types
    bolB, _ := json.Marshal(true)          // "true"
    intB, _ := json.Marshal(1)             // "1"
    fltB, _ := json.Marshal(2.34)          // "2.34"
    strB, _ := json.Marshal("gopher")      // "\"gopher\""

    // Marshal slices and maps
    slcB, _ := json.Marshal([]string{"apple", "peach", "pear"})
    // ["apple","peach","pear"]

    mapB, _ := json.Marshal(map[string]int{"apple": 5, "lettuce": 7})
    // {"apple":5,"lettuce":7}

    // Marshal structs with and without tags
    res1 := &response1{Page: 1, Fruits: []string{"apple", "peach"}}
    res1B, _ := json.Marshal(res1)
    // {"Page":1,"Fruits":["apple","peach"]}

    res2 := &response2{Page: 1, Fruits: []string{"apple", "peach"}}
    res2B, _ := json.Marshal(res2)
    // {"page":1,"fruits":["apple","peach"]}

    // Unmarshal into map[string]interface{}
    byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
    var dat map[string]interface{}
    json.Unmarshal(byt, &dat)
    num := dat["num"].(float64)                // 6.13
    strs := dat["strs"].([]interface{})
    str1 := strs[0].(string)                   // "a"

    // Unmarshal into struct
    str := `{"page": 1, "fruits": ["apple", "peach"]}`
    res := response2{}
    json.Unmarshal([]byte(str), &res)
    fmt.Println(res.Fruits[0])                 // "apple"

    // Stream directly to io.Writer
    enc := json.NewEncoder(os.Stdout)
    enc.Encode(map[string]int{"apple": 5, "lettuce": 7})
}
```

## XML

```go
import "encoding/xml"

type Plant struct {
    XMLName xml.Name `xml:"plant"`
    Id      int      `xml:"id,attr"`    // "attr" makes it an XML attribute
    Name    string   `xml:"name"`
    Origin  []string `xml:"origin"`
}

func TestXMLHandling() {
    coffee := &Plant{Id: 27, Name: "Coffee", Origin: []string{"Ethiopia", "Brazil"}}

    // MarshalIndent for readable output
    out, _ := xml.MarshalIndent(coffee, " ", "  ")
    fmt.Println(string(out))
    // <plant id="27">
    //   <name>Coffee</name>
    //   <origin>Ethiopia</origin>
    //   <origin>Brazil</origin>
    // </plant>

    // Add XML header
    fmt.Println(xml.Header + string(out))

    // Unmarshal
    var p Plant
    xml.Unmarshal(out, &p)
    fmt.Println(p) // {27 Coffee [Ethiopia Brazil]}

    // Nesting with tag directives: parent>child>plant
    type Nesting struct {
        XMLName xml.Name `xml:"nesting"`
        Plants  []*Plant `xml:"parent>child>plant"`
    }
    nesting := &Nesting{Plants: []*Plant{coffee, tomato}}
    out, _ = xml.MarshalIndent(nesting, " ", "  ")
    fmt.Println(string(out))
}
```
