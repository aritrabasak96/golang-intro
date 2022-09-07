package main

import "fmt"

func annoFunc1(){

	// type 1
	// immediately invoked function   
	func(){
      fmt.Println("i am iv function")
	}()

	// type 2
	// function as a variable 
	f := func(){
		fmt.Println("function as a variable")   
	}
	f()

	/* both are same

	 var f func() = func(){
		fmt.Println("function as a variable")
	}
	f()
	*/

	// return value of a annonymous function 
	// myfunc =>  is a variable that is assign to a annonymous function 
	var myfunc func(string,int) (int,int)

	myfunc = func(a string,b int) (int,int){
	   fmt.Println("return0---",a,b)
	   
	   return 12,34
	}

	a,b := myfunc("hello",12)
    fmt.Println(a,b)
}

func main(){

	annoFunc1();
}
