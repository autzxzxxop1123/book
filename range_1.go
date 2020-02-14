package main

import "fmt"

func main(){
	r := []bool{true, false, true, true, false, true}
	for i,v := range r {
		fmt.Printf("%t\n",i, v)
	}
}