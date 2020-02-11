package main

import "fmt"

func main(){

	x := [4]string{"A", "B", "C", "D"}
	fmt.Println(x[:])
	fmt.Println(x[:1])
	fmt.Println(x[1:])
}