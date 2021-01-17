package main 

import ("fmt")

type Creature interface{

	Walk()
	
	
}

type Aritra struct{
	name string
	age int
	habit string
}

type Cat struct{
	size int
	age int
	eat string
}


// you implements the walking method
func (a Aritra) Walk(){
	fmt.Println("aritra can walk slower")
}


// you implements the walking method
func (c Cat) Walk(){
	fmt.Println("cat can walk faster")
}


func myInterface(){
	
	var ae Creature = Aritra{name:"aritra",age:21,habit:"coding"}
	ae.Walk()

	cat := Cat{size:1,age:6,eat:"fish"}
	cat.Walk()

}

func main()  {
	
	myInterface()
}


// --------------- NOTE --------------------
/* 
  
struct -> describe the data 
interface -> describe the behaviour 

*/
