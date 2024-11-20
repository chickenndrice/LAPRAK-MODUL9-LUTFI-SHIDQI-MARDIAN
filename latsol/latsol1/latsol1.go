package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n%2 == 0 {
		n = n / 2
		fmt.Print(n)
	} else {
		n = n / 2 + 1
		fmt.Print(n)

	}
}
