package main

// Sample program to teach the mechanics of escape analysis.

type user struct {
	name  string
	email string
}

func main() {
	u1 := createUserV1()
	u2 := createUserV2()

	println("u1", &u1, "u2", u2)
}

// createUserV1 creates a user value and passed a copy back to the caller
func createUserV1() user {
	u := user{
		name:  "George",
		email: "george@golang.expert",
	}

	println("V1", &u)

	return u
}

func createUserV2() *user {
	u := user{
		name:  "George",
		email: "george@golang.expert",
	}

	println("V2", &u)

	return &u
}
