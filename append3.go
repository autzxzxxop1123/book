package main

import "fmt"

func main(){
	names := []string{"dog","cat"}
	name := []string{"ant"}
	name = append(names, name...)

	fmt.Println(name)
	
}