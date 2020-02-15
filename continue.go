package main

import "fmt"

func main(){
	for i := 0; i < 10; i++{
		fmt.Println("a")
		continue
		fmt.Println("b")
	}
	fmt.Println("c")
}