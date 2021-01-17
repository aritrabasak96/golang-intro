package main

import "fmt"

type student struct{

	name string
	age int
	address string
}

// this function is called method 
// *** defination of method => functions that executes in context of a type ***
func (s student) classRoom(){
	
	new_name := s.name + " Basak"
	fmt.Println("new name is---",new_name)
}

func methodFunc(){

    s := student{
		name: "aritra",
		age:12,
		address:"kolkata",
	} 

	s.classRoom()
}


func main(){
	
	methodFunc()
}