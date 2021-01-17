// there are many ways where you can declare varibales in go

// 1)
package main

import ("fmt"
		"strconv"
)


func myfun(){

   // create variable now, declare later
	var i int;
	i = 10;

	// create variable and declare now
	var m int = 10;
	var j string = "aritra";
	var p float32 = 12.34;
	var b bool = true; // if you do not assign any value then by default go will set the value to false

	// only applicable in inscope function variable
	// you do not have many control if you declare this variable like this
	n := 12;

	// another way to create varibales
	var (

		name string = "basak";
		number float64 = 234.55;
	)

	fmt.Printf("%v,%T",name,name); // %v => variable %T => type of variable

}


// convert variable from one format to another
func convertVar(){

	var i int = 97;
	var j string;

	//  ----- convert int to float
	//  j = float32(i)

	// ----- convert int to string
	// j = string(i); // this will give you unpredictive result, because go convert this as ASCII value

	// to convert any value to string properly
	j = strconv.Itoa(i)

	fmt.Printf("%v, %T",j,j);
}

func main()  {

   myfun();
   convertVar();
}


/*
  ----------------------------------- NOTE ------------------------------------

 1) if you create or declare a variable but never used in your function scope then go compiler will give
  you an error.

 2) If you declare a variable outside of any function then it is called package level variable
	if the variable starts from lower case letter then the variable can be accessed by package
	level only
	ex: var name string = ""

	if the variable starts from Upper case letter  then the  variable can be accessed by globally
	ex: var NAME string = ""

*/
