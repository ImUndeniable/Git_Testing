package main

import (
	"fmt"
)


func main(){
	var x = "Wassap!"
	Hello(x)

	var a int = 5
	var b int = 6
	var result = Add(a,b)
	fmt.Println(result)
	

}


//calling function
func Hello(x string){
	fmt.Println("Yooo",x)
}

// return function
func Add(a int ,  b int) int{
	var result int = a + b
	return result
} 