# Structs and Methods

## Structs

Go does not support classes. Instead, it has structs. Structs are collections of fields that allow you to group data together.

Let's make a struct to store the data of Contacts:

```go
type Contact struct {
  name string
  age  int
} 
```

Our Contact struct has two fields, a string and an integer. Now, we can create a new Contact using the following syntax:

```go
x := Contact{"James", 42}
```

x is now a struct object that is initialized with the data provided in the curly braces.

We can also provide the names of the fields when creating a new struct. This makes it easier to read the code. For example:

```go
x := Contact{name: "James", age: 42}
```

We can access the struct fields using the name of the struct and a dot:

```go
x := Contact{"James", 42}
x.age = 8
fmt.Println(x.age)
fmt.Println(x.name)
```

Note that we are able to change the values of the fields by simply assigning them new values.

## Pointers to Structs

Similar to simple pointers, we can also make pointers to structs using the `&` operator:

```go
x := Contact{"James", 42}
p := &x
```

Pointers to structs are automatically dereferenced, meaning we can access the field values by simply using a dot:

```go
x := Contact{"James", 42}
p := &x
fmt.Println(p.age)
```

We could use `(*p).age` to access the age field of the struct, but that looks complicated and hard to read. Go allows you to shorten that syntax and simply use `p.age` instead. We can also use pointers when creating a new struct:

```go
p := &Contact{"John", 15}
fmt.Println(p.name)
```

Now `p` is a pointer to the newly created struct. Pointers to structs are useful, as they allow you to pass them to functions as arguments.

## Methods

We can add functionality to our structs using methods. Methods are simply functions with a special receiver argument. Let's have a look at an example:

```go
func (x Contact) welcome() {
  fmt.Println(x.name)
  fmt.Println(x.age)
}
```

The receiver appears between the func keyword and the method name. In the example above, the receiver is the Contact struct. Note that we can access the receiver struct's fields in the method.

```go
type Contact struct {
    name string
    age int
}

func (contact Contact) show() {
    fmt.Println(contact.name)
    fmt.Println(contact.age)
}
func (ptr *Contact) increase(val int) {
    ptr.age += val
}

func main() {
    contact := Contact{"James", 42}
    contact.increase(8)
    contact.show()
}
```

Since methods are just functions with a receiver argument, we can achieve the same functionality using a regular function that takes a struct as its argument.

In case we need to change the data of the struct in a method, we can use pointers as method receivers. Go automatically dereferences the pointers, so we simply call the method with the struct name, just as we did with a simple receiver. **Since methods often need to modify their receiver, pointer receivers are more common than value receivers**.

We can also use type on the existing datatypes and attach receiver function on them.

```go
type Cards []string
```

We can use this cards anywhere like `cards := Cards{}`.

## Embedded Structs

Golang doesn't have the concepts of objects and inheritance. The design of the language is strongly opinionated and follows object-oriented principles very closely, hence it favors **composition over inheritance**. Golang uses structs and methods to emulate objects.

Golang allows you to compose structs inside of other structs. There are two ways of composing structs: **anonymous embedded fields** and **explicit fields**.

### Anonymous Embedded Fields

An anonymous embedded field is a struct field that doesn't have an explicit field name. Structs that have an anonymous field obtain all of the methods and properties of the nested struct. These methods and properties are called **"promoted"** methods and properties. Anonymous embedded fields can also be accessed by the type name directly.

The nuance comes from when structs declare methods and properties that also exist on the embedded struct. In this case, the methods and properties from the embedded struct are "shadowed" and the methods and properties on the struct you are calling will be used. You can still access the shadowed methods and properties by specifically accessing the embedded field and then calling the method or property.

```go
type Bird struct {}

func (b Bird) makeSound() {
  fmt.Println("chirp")
}

type Eagle struct {
  Bird // anonymous embedded field
}

func main() {
    e := Eagle{name: "Baldie"}
    e.makeSound() // chirp
    e.Bird.makeSound() // chirp
}
```

### Explicit Fields

You can also explicitly name a field to avoid confusion with promoted and shadowed fields.
Here, in order to call Bird's `makeSound()` method, one has to explicitly access the nested Bird object using `.b`. It is no longer possible to access the underlying `.Bird` field directly.

```go
type Eagle struct {
  b Bird // explicit field
}

func (e Eagle) makeSound(){
  fmt.Println("caw")
}

func main() {
    e := Eagle{name: "Baldie"}
    e.makeSound() // caw
    e.b.makeSound() // chirp
    e.Bird.makeSound() // error
}
```

In production code, it's often better to use explicit fields over anonymous embedded fields to make the code more readable and not surprise yourself with shadowed fields.

## Maps vs. Structs

### Map

- All key and value are of same type.
- When keys are indexed we can iterate over them.
- Closely related and significant value type.
- Don't need to know all the keys at compile time.
- Keys are indexed - we can iterate over them.
- Reference type.
- Zero value for a map is empty map.

### Struct

- All values can be of different type.
- Need to know all the different fields at compile time.
- Keys don't support indexing.
- Value type.
- Zero value for struct will be according to the struct field type.

### When to use?

**When to use structs?** If we have a close set of keys, meaning fixed data size with keys, we will be using structs. Using structs is a safe and easy way while working with JSON data also.

**When to use maps?** If we are creating some kind of relationship between keys and values and we don't really know what that collection of values is going to be at compile time or as we are writing our code, then we have a great use-case for using a map.

In most cases, the vast majority of time, we mostly use structs rather than maps. But it all depends on the nature and type of the application and the requirement of the project.

**Note**: Small case means that is a private field or member, Capital case means that is a public field or member.
