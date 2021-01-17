package main 

import ("fmt"
)

func myMap(){

	// map is a data structure that store data as a key value pair in golang 
	
	// syntax to create a map 
	// map[data type of the key]data type of the value
	// ex: map[string]int => key is string, value is int  

	// ****** map does not gurantee the order when you retrive it ********

	m1 := map[string]int {

		 "ram" : 23,
		 "shyam" : 24,
		 "ramesh" : 21,
	}

	m2 := map[int]string {

		1 : "aritra",
		2: "jodu",
		3 : "modhu",
	}

	// create map and declare data later 
	m3 := make(map[int]int); // create the map data structure 
	m3 = map[int]int{  // declare data later 
		1 : 100,
		2: 342,
	}
	
	// add value later 
	m3[3] = 234;
	
	// delete data from map 
	fmt.Println(m1,m2,m3)

	delete(m3,3)  // the map variable and the key you want to delete 

	fmt.Println("new m3---------------",m3)

	
}

func main(){

	myMap()
}