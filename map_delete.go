package main 

import "fmt"

func main(){
	elements := make(map[string]string)
	elements["D"] = "DOG"
	elements["C"] = "CAT"
	elements["A"] = "ANT"
	fmt.Println(elements)

	delete(elements, "C")
	fmt.Println(elements)

}