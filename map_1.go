package main 

import "fmt"

func main(){
	elements := make(map[string]string)
	elements["D"] = "DOG"
	elements["C"] = "CAT"
	elements["A"] = "ANT"

	h := elements["D"]
	fmt.Println(h)

	n, found := elements["S"]
	fmt.Println(n)
	fmt.Println(found)
}
