package main 

import ("fmt")

func myFor(){
	
	// no first brackets  

	for i := 0; i<10 ; i++{

		fmt.Println(i)
	}

	// nested loop 
	for p := 0; p<10 ;p++{

		for t := 0; t<12 ; t++{

			fmt.Println(p,t)
		}
	}

	// another way of writing for loop 
	for n,m := 0,0; n < 5; n,m = n +1, m+1{

		fmt.Println(n,m)  // o/p => n = 0,1,2,3,4    m = 0,1,2,3,4
	} 
	 
	var s string = "new name";
	
	for k,v := range s{
		fmt.Println("k---",k,string(v))
	}

	// loop a array or slice 
	loop := []int{2,4,9,12}
	
	for num,key := range loop{
		fmt.Println("num---",num,key)
	}
}

func main(){

	myFor()
}