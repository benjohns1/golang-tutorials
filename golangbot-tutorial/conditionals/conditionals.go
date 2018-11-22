package main

import "fmt"

func main() {
	// if/else
	if num := 10; num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// for
	for i := 1; i <= 10; i++ {
		fmt.Printf(" %d", i)
	}
	fmt.Println()

	for i := 0; i < 10; i++ {
		if i > 5 {
			break
		}
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\nline after for loop")

	for no, i := 10, 1; i <= 10 && no <= 19; i, no = i+1, no+1 {
		fmt.Printf("%d * %d = %d\n", no, i, no*i)
	}

	// switch
	count := 6
	switch count {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	case 4:
		fmt.Println("4")
	case 5, 6, 7:
		fmt.Println("5")
	default:
		fmt.Println("6")
	}

	// expressionless switch
	num := 150
	switch {
	case num > 0 && num <= 50:
		fmt.Println("num is greater than 0 and less than 50")
	case num >= 51 && num <= 100:
		fmt.Println("num is greater than 51 and less than 100")
	case num >= 101 && num <= 200:
		fmt.Println("num is greater than 100 and less than 200")
		fallthrough
	default:
		fmt.Println("num is less than 0 or very large")
	}
}
