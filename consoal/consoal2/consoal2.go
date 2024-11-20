package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n >= 0 {
		fmt.Print("Bilangan Positif")
	} else {
		fmt.Print("Bilangan Negatif")
	}
}