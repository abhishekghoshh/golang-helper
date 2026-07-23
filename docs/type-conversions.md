# Type Conversions, Casting, and Type Assertions

Interfaces in Go provide a way to specify the behavior of an object: if something can do this, then it can be used here.
Interfaces are a big deal in Go. If a variable's type is that of an interface, then you can be confident that the object referenced by the variable implements the methods prescribed by the interface.

```go
type Stringer interface {
    String() string
}
var x Stringer
```

The only guarantee an interface provides is that the object it points to will implement its prescribed methods. Nothing more.

That is also the reason why the empty interface (`interface{}`) is almost useless because it doesn't guarantee anything! Its primary usefulness is only due to the lack of Generics.

## Type Assertions

Although the variable is an interface, it will still reference an object that contains fields and other methods (exported or unexported).
To access those fields and methods, you need to type assert. Type assertion basically proclaims that the object is something else (either another interface or struct).

```go
type A struct {
    name string
}

// A implements Stringer interface
func (a A) String() string {
    return "Hello"
}

func main() {
    var x Stringer
    x = A{name: "sam"}
 
    fmt.Println(x.String()) // Output: Hello
    fmt.Println(x.(A).name) // Output: sam
}
```

You can type assert by using this syntax: `x.(A)`, where x is the (interface) variable and (A) is the type you are proclaiming x to really being.

```go
fmt.Println(x.name) 
// Error: x.name undefined (type Stringer has no field or method name)
```

If you try and access name without type asserting, then you will get a compile-time error.

Of course, if you type assert incorrectly (assert that x is something it's not), the application will crash with a run-time panic. To avoid this, you can use a **type switch** or the **comma, ok idiom**.

## Type Conversions

In Go, all variable types are distinct from one another — even if behind the scenes, they are stored with exactly the same structure in memory or are aliases of each other.

This means an int type is distinct from an int64 type, even though on a 64-bit machine they are stored in memory the same way.

Type conversion is required when you need to convert one variable type into another, usually because a particular variable or function argument requires it. If you want to convert an int64 to an int, the syntax is: `x := int(y)`, where y is an int64 and x is an int.

Type conversion will create a copy of the original data, so it's not necessarily a "free" operation.

Go's type conversion will do its best to maintain the same value in the new data type if possible. To do that, it may transform the underlying bit structure.

It can also fail to convert accurately on some occasions: converting from a larger to smaller data type, or from a signed to an unsigned, or from a large int64 to a float64 are common culprits.

## Casting

Casting is seldom used in Go. Even advanced developers will rarely (if ever) see explicit casting syntax in their code.

When coupled with the unsafe package, there is one elegant use-case of casting. It is potentially dangerous, and there are safer ways to achieve the same objective. It's when you want to convert between 2 different types of structs which have precisely the same underlying data structure.

Type Conversion doesn't work for this scenario, but this does:

```go
import "unsafe"

type Z struct {
    A int
    B string
}

type Y struct {
    C int
    D string
}

func main() {
    z := Z{A: 1, B: "sam"}
    y := *(*Y)(unsafe.Pointer(&z))
}
```

Attempting to convert z (of type Z) to type Y will not work with the conversion syntax `y := Y(z)`. In this example, casting is required to bypass Go's type safety checks.
