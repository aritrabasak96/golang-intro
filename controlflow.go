package main 

import "fmt"

// where defer is used 
// mostly network call, when you want to terminate the request at the last, when the function complete 
// it's execution 
func myDefer(){
	
	// 'defer' will run the command after the function ends 
	// o/p => second third first
	defer fmt.Println("first")
	fmt.Println("second");
	fmt.Println("third");

	// this time fmt.Println() will call when the function ends, so first time 1 then 3 then 4 and then 2
	// but the value of a will not change  
	// o/p => second 15 , first 10 

	a := 10 // 1
	defer fmt.Println("first a",a); // 2
	a = 15 // 3
	fmt.Println("second a",a) // 4

}

func main(){
  
	myDefer()
}