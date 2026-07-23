# Controllers (HTTP Routing)

Three approaches to building HTTP APIs in Go.

## Shared Service Layer

```go
type Person struct {
    Id        int    `json:"id" xml:"id"`
    FirstName string `json:"firstName" xml:"firstName"`
    LastName  string `json:"lastName" xml:"lastName"`
    Age       int    `json:"age" xml:"age"`
    Gender    string `json:"gender" xml:"gender"`
}

var persons = []Person{
    {Id: 1, FirstName: "Abhishek", LastName: "Ghosh", Age: 25, Gender: "Male"},
    {Id: 2, FirstName: "Nasim", LastName: "Molla", Age: 27, Gender: "Male"},
}

func GetPersons() []Person    { return persons }
func GetPerson(id int) Person { /* find by id */ return Person{} }
func AddPerson(p *Person)     { persons = append(persons, *p) }

// Content-Type-aware response writer
func AddToResponse(w http.ResponseWriter, r *http.Request, data interface{}) {
    switch r.Header.Get("Content-Type") {
    case "application/xml":
        w.Header().Add("Content-Type", "application/xml")
        xml.NewEncoder(w).Encode(data)
    case "application/json":
        w.Header().Add("Content-Type", "application/json")
        json.NewEncoder(w).Encode(data)
    default:
        t := reflect.TypeOf(data)
        if t.Kind() == reflect.Struct || t.Kind() == reflect.Slice {
            w.Header().Add("Content-Type", "application/json")
            json.NewEncoder(w).Encode(data)
        } else {
            fmt.Fprintf(w, data.(string))
        }
    }
}

// Signal handling for graceful shutdown
func cleanup() {
    c := make(chan os.Signal, 2)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    go func() {
        <-c
        fmt.Println("Ctrl+C pressed - saving data...")
        savePersonsToDisk()
        os.Exit(0)
    }()
}
```

## 1. Basic HTTP Controller (Default ServeMux)

```go
func HttpBasicController() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        AddToResponse(w, r, "Hi this is http basic controller")
    })
    http.HandleFunc("/persons", func(w http.ResponseWriter, r *http.Request) {
        AddToResponse(w, r, GetPersons())
    })
    log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
```

## 2. HTTP Mux Controller (Explicit ServeMux)

Uses `http.NewServeMux()` instead of the default.

```go
func HttpMuxController() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        AddToResponse(w, r, "Hi this is http mux controller")
    })
    mux.HandleFunc("/persons", func(w http.ResponseWriter, r *http.Request) {
        AddToResponse(w, r, GetPersons())
    })
    log.Fatal(http.ListenAndServe("localhost:8080", mux))
}
```

## 3. Gorilla Mux Controller (Advanced Routing)

Uses `github.com/gorilla/mux` for URL path variables, method-based routing.

```go
func GorillaMuxController() {
    router := mux.NewRouter()

    router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        AddToResponse(w, r, "Hi this is Gorilla mux controller")
    })

    router.HandleFunc("/persons", func(w http.ResponseWriter, r *http.Request) {
        AddToResponse(w, r, GetPersons())
    }).Methods("GET")

    // Path variable with regex constraint
    router.HandleFunc("/persons/{personId:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        personId, _ := strconv.Atoi(vars["personId"])
        AddToResponse(w, r, GetPerson(personId))
    }).Methods("GET")

    // POST with JSON body decode
    router.HandleFunc("/persons", func(w http.ResponseWriter, r *http.Request) {
        var p Person
        json.NewDecoder(r.Body).Decode(&p)
        AddPerson(&p)
        AddToResponse(w, r, p)
    }).Methods("POST")

    log.Fatal(http.ListenAndServe("localhost:8080", router))
}
```

**Features:** Persistent storage via JSON file, content-type negotiation (JSON/XML), and graceful shutdown with signal handling.
