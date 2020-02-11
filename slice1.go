package main	

import "fmt"

func main(){

	text := [4]string {"A", "B", "C", "D"}
	fmt.Println(text)

	x := text[0:2]
	fmt.Println(x)

	y := x[2:4]
	fmt.Println(y)

	z := y[0:1]
	fmt.Println(z)

	z[0] = "X"
	fmt.Println(text)
}