package main

import(
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}


	switch time.Now().Weekday() {
		case time.Saturday,time.Sunday:
			fmt.Println("Hurrah! its weekend")
		default:
			fmt.Println("Its a weekday")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I am a bool")
		case int:
			fmt.Println("I am an integer")
		default:
			fmt.Println("Dont know the type ",t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}

