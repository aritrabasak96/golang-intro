package main

import ("fmt")

func myPanic(){

	var a int  = 10
	var b int = 10;

	var diff int = a - b;

	if diff == 0{

		// if you want to throw an error
		// although go does not have any try catch systax
		// **** panic happend after execution of 'defer'

		fmt.Println("This statement will execute first");


    // recover() => return nil if the application is not panicking
		//           => return an error if the application start panicking

		// anonnymous function
		defer func()  {
			if err := recover(); err != nil{
				fmt.Println("This statement will execute---",err)
			}
		}()

		fmt.Println("before panicking")

		panic("something went wrong"); // throw a panic



		fmt.Println("This statement will not execute");

		// o/p => execute from start....

		// 1) 'This statement will execute first' get called
		// 2) then the defer function will execute at last
		// 3) 'before panicking' get called
		// 4) then the application start panicking
		// 5) so the last statement will not execute
		// 6) then the defer function will start execute
		// 6.1) you call recover() function which handle the error and do not throw a bad error
		// 6.2) 'This statement will execute---something went wrong' will be printed

	} else{
		fmt.Println("pass")
	}
}

func main(){

	myPanic()
	fmt.Println("main function") // printed at last 
}
