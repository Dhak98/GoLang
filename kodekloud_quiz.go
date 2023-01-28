package main

import "fmt"

func main() {
    var a,b int
	fmt.Println("Enter the values for A & B")
	fmt.Scan(&a,&b)
	var c int = a + b
	if c == 50 {
		fmt.Println("It is 50")
	}else{
		fmt.Println("Tata, Bye Bye")
	}

}