package main 

import ("fmt")


type student struct{
	
	name string
	age uint 
	degree []string
}

// we do not have any database connection right now 
// that is why we are creating our own array data type 
type studentHandler struct{
	
	// my key is a string and data is student 
	store map[string]student
}

func (st *studentHandler) putName(){

}

func main(){

}