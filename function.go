package main 

import ("fmt")

// passing value as a parameter 
func paramFunc(name string){

	fmt.Println("hello "+name+" how are you");
}

// create multiple value unconditionally 
func mulFunc(name,skill string){
	
	fmt.Println("first time---",name+" "+skill);

	// assign a new value 
	name = "banti"
	// ---------------- NOTE ---------------------

	// go compiler will create a new copy of the variable and assign a new value 
	// so now in your program you have two variable with a same name 
	// one is, name = "banti" which is in the scope of mulFunc()
	// and second is, name = "aritra" which is in the scope of main() function  

	// -----------------  unconditionally, you waste some memory -------------------

	// ------------- solution is ----------------
	// you should pass the pointer as a variable 

	/* 
	
       func main(){

      	paramFunc("aritra")

	    name := "aritra";
	    skill := "programming";
	    mulFunc(&name,&skill); // pass the memory address 
	 }
	 
	 func mulFunc(name,skill *string){
		  
		fmt.Println("first time---",*name+" "+*skill);
		*name = "banti"  // then banti will refelect in both function scope 
	 }

	*/


	fmt.Println("second time---",name+" "+skill);
}

// pass multiple value 
func mulParamsFunc(val string,num ...int)  {   // func mulParamsFunc(num ...int,val string) // this is wrong
	
	fmt.Println("num---",num)
	fmt.Println("val-----",val)
}

func returnFunc() int {
	
	val := 10
	return val;

	// -------------- you can also return a pointer 

// func returnFunc() *int {
	
// 	val : = 10
// 	return &val;
	
   // --------------- NOTE ----------------
   // when you return the memory address, go compiler actually store the variable in the memory heap 
   // although when the function ends, all variables are destroyed, but if you pass memory address then go compiler will do this for you 


// }

}

func secondParamsFunc(a,b float64) (float64,error){

	if b == 0.0{
		return 0.0, fmt.Errorf("new error")
	}

	return a / b, nil
}


func main(){

	paramFunc("aritra");

	mulParamsFunc("all numbers",1,2,3,4,5,6);

	s := returnFunc()	
	fmt.Println("s---",s)

    a,b := secondParamsFunc(5.0,0.0);   
	
	fmt.Println("a----",a)
	fmt.Println("b----",b)
	
 

	name := "aritra";
	skill := "programming";
	mulFunc(name,skill);
}