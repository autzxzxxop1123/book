package main

import "fmt"

func main(){
	elements := make(map[string]map[string]string)
	elements["A"] = map[string]string{"name": "ant", "name_1" :  "dog"}
	elements["He"] = map[string]string{"name": "cat", "name_1": "rat"}
	fmt.Println(elements)
	fmt.Println(elements["He"]["name_1"])

}