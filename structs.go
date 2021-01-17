package main

import ("fmt")

// structure is a way to store group of variables ex: if you want to store a doctor information
// - then you have to store it's name, id , age and more


// when we create a structure we create a data type

// define a datatype whose name is Doctor
type Student struct {

	name string
	age int
	organization string
	degree []string // slice

}

func myStruct(){

	student := Student{name:"Rahul",age:30,organization:"Medical",degree: []string{"b.tech","m.tech","phd"},}

	fmt.Println("name----",student.degree);

}

// another way of assigning values 
func my1Struct(){

	student := Student{};
	student.name = "aritra";
	student.age = 12

	fmt.Println("student1---",student)
}

// create a different type of struct format
func my2Struct(){
	
	teacher := struct{name string; age int}{name:"Doyel",age:23}

	fmt.Println("teacher-----",teacher);  // o/p => {"Doyel",23}
}


func main(){

	myStruct()
	my1Struct()
	my2Struct()
	
}
