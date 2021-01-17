package main 

import "fmt"

// go does not support object oriented programming pattern 

// but go support composition 

type Animal struct{

	blood string
	walk string

}

type Human struct{
   
    // ***** it is not traditional inheritence, you can not say that the human is a animal 
    // ***** you just embed animal to human  	
	Animal // composite (or you can say inherit)
	dance string 

}


func main(){

	man := Human{}  // create instance
	man.blood = "red"
	man.walk = "fast";
	man.dance = "hip hop"

	fmt.Println("animal-----",man);  // o/p => {{red,fast}, hip hop}
}