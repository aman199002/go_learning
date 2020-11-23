package main
import "fmt"

func main() {
	n := 5
	if n%2 == 0 {
		fmt.Println(n, " is even")
	} else {
		fmt.Println(n," is odd")
	}
	if i := 9; i < 0 {
		fmt.Println(i," is negative")
	} else if i < 10 {
		fmt.Println(i," has 1 digit")
	} else {
		fmt.Println(i," has 2 digits")
	}
}
