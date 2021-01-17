package main 

import ("fmt")

type student struct{
	name string
	age uint
	degree []string
}

func (s1 *student) assignName(nm string) *student{
	
	// 4) s1 points the same object in the memory >> s1 == s 
	// &s1 => 0xc00006c330 and &s => 0xc00006c330
	s1.name = nm 
	return &student{name:s1.name,age:s1.age,degree:s1.degree}
}

func main(){
	
	// 1) create an object or instance variable
	s := student{}
	// s takes a memory address in ram >> 0xc00006c330
	s.name = "aritra"; 
	s.age = 21;
	s.degree = []string{"java","golang","python"}

	// 2) if you want to change the name, you can 
	s.name = "banti"
	 
	// 3) call the method 
	val := s.assignName("ab")
	

}