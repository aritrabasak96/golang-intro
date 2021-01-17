package main

import ("fmt")

func constant(){

	 const a int = 10;
	 //a = 14; // you can not assign . You get an error during run time
	 fmt.Println(a)
}

// iota 
// Enumerated constants
func myiota(){

	const (
		
		// compiler will automatically assign values to other constants
		// _ = 0 + 23.11 (iota value = 0)
		// first = 23.11 + 1 = > 24.11
		// second = 24.11 + 1 => 25.11 
		_ = iota + 23.11
		first
		second
		third
		fourth
	)

	fmt.Println(second)
}



func  main()  {
	 
    myiota()	
	constant()
}