package main

func main() {
	// Declare a variable of type int with a value of 10
	count := 10

	// Display the "value of" and "address of" count.
	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

	increment(count)

	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")
}

/*
Increment declare count as a pointer variable whose value is always an
address and points to values of type int.
*/
func increment(inc int) {
	// Increment the "value of" inc.
	inc++
	println("inc:\tValue Of[", inc, "]\tAddr Of[", &inc, "]")
}

/*
Goroutines works in a copy of this element, with doesn't affect the original element
see by the address values in the print execution, the value address of inc is different of
the original, the modification occurs in isolation, this is named "Value Semantics"
*/
