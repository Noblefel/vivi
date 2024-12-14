Some helpers to get user input from terminal. Thanks to [keyboard](https://github.com/atomicgo/keyboard) from atomicgo

```go
go get github.com/Noblefel/vivi
```

### Example: select from multiple choice

```go
func main() {
	// basic
	fmt.Println("are you ready to play?")
	ready := vivi.Choices("Yes", "No")

	if ready == 1 {
		return
	}

	// slice
	fmt.Println("----------------")
	fmt.Println("which you like the most ?")

	foods := []string{
		"Burger 🍔",
		"Pizza 🍕",
		"Sushi 🍣",
		"Steak 🥩",
		"Spaghetti 🍝",
		"Fries 🍟",
	}

	fav := vivi.Choices(foods...)
	fmt.Println("me too like", foods[fav], "😋")
}

```

### Example: passwords

```go
func main() {
	fmt.Print("enter password please: ")
	pw := vivi.Password("*")

	if pw != "abc" {
		fmt.Println("wrong!!!")
		return
	}
}
```

### Example: hidden input

```go
func main() {
	fmt.Print("say something very secret: ")
	secret := vivi.Password("")
	fmt.Println("this secret is safe")
}
```
