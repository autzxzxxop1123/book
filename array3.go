package main

import "fmt"

func main(){
	number := [2][3][2]int{
		{
			{1,2},
			{10,20},
			{100,200},
		},
		{
			{8,9},
			{80,90},
			{800,900},
		},

	}
	fmt.Println(number)
	fmt.Println(number [1][2][0])
}