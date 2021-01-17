// array
package main

import "fmt"

func  myArray()  {

	// Differemt way to create and declare an array
	// 1)
	array1 :=  [3]int{12,34,45}  // [...] => i can push any number just grab it

	// 2)
	var array2 [3]string; // can accept 3 string values

	array2[0] = "aritra";
	array2[1] = "basak"

	// 3)
	var array3 [2]int = [2]int{12,22} // [...]int{12,22}

	// 4) 2D array
	var array4 [3][3]int = [3][3]int{[3]int{1,2,3},[3]int{3,5,7},[3]int{1,7,2}}

	// 5) copy from one array to another
	array5 := [4]int{2,4,6,8}

	temparray := array5;   // temparray does not store the reference or the memory ref to array5
	                       // insted it will create a new array

	temparray[2] = 12 // assign a new value

	// array5 => [2,4,6,8]
	// temparray => [2,3,12,8]

	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array3)
	fmt.Println(array4)
	fmt.Println("array5",array5)
	fmt.Println("temparray",temparray);

}

// slice
func mySlice(){
	// everything that we can do with array, we can do with slice
	// unlike a array, slice do not have a fixed size
	a := []int{2,3,5,6,8,9}

  a = append(a,12,23,45,66)

	fmt.Println("slice",a)

	b := [...]int{3,6,10,33,12}

	c := b[2:] // from 2 to end o/p =>

	fmt.Println("slice--c",c)
	fmt.Println("slice----b",b)
}

func checkMySlice(){

	a := []int{2,4,5,6,7}

	b := append(a[2:])

	fmt.Println("a-----",&a);
	fmt.Println("b-----",&b);
}

func main()  {

	mySlice();
	myArray();
	checkMySlice();
}

// --------------------------------- NOTE -----------------------------
//  for array and struct
//  a := [2]int{23,45}
//  b := a // you copy the value to b
//  if you change any value e.x a[1] = 100
//  then only a will change , not b

// -----------------------------------------
// for slice and map
// if you change any value e.x a[1] = 120
// then both a and b will change 
