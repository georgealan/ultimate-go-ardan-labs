package main

func main() {
	// Declare variable of type int with a value of 10
	count := 10

	// Display the "Value Of" and "Address Of" count.
	println("count:\tValue Of[", count, "]\t\tAddr Of[", &count, "]")

	// Pass the "Address Of" count.
	increment(&count)

	println("count:\tValue Of[", count, "]\t\tAddr Of[", &count, "]")
}

func increment(inc *int) {
	// Increment the "value of" inc.
	*inc++
	println("inc:\tValue Of[", inc, "]\tAddr Of[", &inc, "]\tValue Points To[", *inc, "]")
}

/*
Now with this approach we receive in function params the address of the integer value otherwise the
value. With that the value returned will be modified. Goroutine have indirect access to element in memory
this is called pointer semantics.
*/
