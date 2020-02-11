package main

import "fmt"
// 5 = ความยาว ไม่ระบุความจุ ความจุเลยเท่ากับความยาว คือ 5 
func main(){
	x := make([]int , 5)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
// 5= ความยาว 10 = ความจุ
	y := make([]int, 5, 10)
	fmt.Println(y)
	fmt.Println(len(y))
	fmt.Println(cap(y))
}