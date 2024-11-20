package main

import "fmt"

func main() {
	var x, y int
	var XY, YX bool
	fmt.Scan(&x, &y)
	if y%x == 0 {
		XY = true
	}	
	if x%y == 0 {
		YX = true
	}
	fmt.Println(XY)
	fmt.Print(YX)
}
