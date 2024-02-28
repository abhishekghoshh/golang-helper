package introduction

import "fmt"

/*
https://gobyexample.com/structs
Structs are mutable.

Access struct fields with a dot. You can also use dots with struct pointers - the pointers are automatically dereferenced.
To access the field X of a struct when we have the struct pointer p we could write (*p).X.
However, that notation is cumbersome, so the language permits us instead to write just p.X, without the explicit dereference.


Go does not have classes. However, you can define methods on types.
A method is a function with a special receiver argument.
The receiver appears in its own argument list between the func keyword and the method name.

	type Vertex struct {
		X, Y float64
	}

	func (v Vertex) Abs() float64 {
		return math.Sqrt(v.X*v.X + v.Y*v.Y)
	}
Remember: a method is just a function with a receiver argument.

You can declare methods with pointer receivers.
This means the receiver type has the literal syntax *T for some type T. (Also, T cannot itself be a pointer such as *int.)
For example, the Scale method here is defined on *Vertex.
Methods with pointer receivers can modify the value to which the receiver points (as Scale does here).
Since methods often need to modify their receiver, pointer receivers are more common than value receivers.
	type Vertex struct {
		X, Y float64
	}

	func (v Vertex) Abs() float64 {
		return math.Sqrt(v.X*v.X + v.Y*v.Y)
	}

	func (v *Vertex) Scale(f float64) {
		v.X = v.X * f
		v.Y = v.Y * f
	}

	func main() {
		v := Vertex{3, 4}
		v.Scale(10)
		fmt.Println(v.Abs())
	}
Try removing the * from the declaration of the Scale function on line 16 and observe how the program's behavior changes.
With a value receiver, the Scale method operates on a copy of the original Vertex value. (This is the same behavior as for any other function argument.)
The Scale method must have a pointer receiver to change the Vertex value declared in the main function.
For the statement v.Scale(5), even though v is a value and not a pointer, the method with the pointer receiver is called automatically.
That is, as a convenience, Go interprets the statement v.Scale(5) as (&v).Scale(5) since the Scale method has a pointer receiver.

There are two reasons to use a pointer receiver. As types are passed by value not reffrence
The first is so that the method can modify the value that its receiver points to.
The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct,


Methods can be defined for either pointer or value receiver types.
Go automatically handles conversion between values and pointers for method calls.
You may want to use a pointer receiver type to avoid copying on method calls or to allow the method to mutate the receiving struct.
*/

func Structs() {
	abhishek := Person{
		"Abhishek",
		25,
		// contact info will have the default value
		contactInfo{},
	}
	kushal := Person{
		name: "Kushal",
		age:  25,
		// contact info will have the default value
	}
	fmt.Println(abhishek, kushal)
	abhishek.welcome()
	abhishek.increase(2)
	abhishek.showAge()

	nasim := Person{
		name: "Nasim",
		age:  20,
		contactInfo: contactInfo{
			email:   "nasim@test.com",
			zipCode: 742101,
		},
	}
	fmt.Println(nasim)
	setName(nasim, "Nasim molla")
	fmt.Println(nasim)

	// If a struct type is only used for a single value,
	// we don't have to give it a name. The value can have an anonymous struct type.
	// This technique is commonly used for table-driven tests.

	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}

type Person struct {
	name        string
	age         int
	contactInfo // this declaration is same as contactInfo contactInfo, we can use this shortcut
}

type contactInfo struct {
	email   string
	zipCode int
}

// passing by value
func setName(p Person, name string) {
	p.name = name
}

// Go supports methods defined on struct types.
func (p Person) welcome() {
	fmt.Println("Welcome", p.name)
}

// Methods can be defined for either pointer or value receiver types.
func (p *Person) increase(age int) {
	p.age += age
}

func (p Person) showAge() {
	fmt.Println("age of", p.name, "is", p.age)
}
