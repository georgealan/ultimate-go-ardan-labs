package main

import "fmt"

// Example represents a type with different fields.
type example struct {
	flag    bool
	counter int16
	pi      float32
}

/*
A technic to prevent issues with heavy use of memory because of the paddings in the middle
and in the end is to set the variables with most big bits in the beginning in descending order.
You don't need use this always, only if you begin have issues with heavy use of memory in your
application.
*/
type exampleTwo struct {
	pi      float32 // Most big bits first followed by descending others.
	counter int16
	flag    bool
}

// Declare a variable of an anonymous type set to its zero value
var e2 struct {
	flag    bool
	counter int16
	pi      float32
}

// Same structures here below, they are identical but for Go they aren't the same.
type george struct {
	flag    bool
	counter int16
	pi      float32
}

type vera struct {
	flag    bool
	counter int16
	pi      float32
}

func main() {
	// Declare a variable of type example set to its zero value.
	var e1 example

	fmt.Printf("Anonymous type variable struct: %+v\n", e2)

	// Declare a variable of an anonymous type and using a struct literal
	e3 := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		flag:    true,
		counter: 10,
		pi:      3.14159,
	}

	fmt.Printf("Anonymous type variable struct literal: %+v\n", e3)
	fmt.Println("flag: ", e3.flag)
	fmt.Println("counter: ", e3.counter)
	fmt.Println("pi: ", e3.pi)

	// Display the value
	fmt.Printf("%+v\n", e1)

	// Declare a variable of type exampleTwo and init using a struct literal.
	e2 := exampleTwo{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	// Display the field values
	fmt.Println("flag", e2.flag)
	fmt.Println("counter", e2.counter)
	fmt.Println("pi", e2.pi)

	// Test same structs
	var a george
	var b vera
	// a = b // This won't run because compiler won't allow assigning named types to another
	// For allow this you need to use conversion like this:
	// a = george(b) // conversion runs OK without compile error
	/*
		But if we use a literal type it will work too, because the compiler allow assigning
		literal structs to named structs
	*/
	a = e3
	fmt.Println(a, b)
}
