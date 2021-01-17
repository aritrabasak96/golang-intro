package main 

import ("fmt"
       "encoding/json")
		

type Book struct{
	 
	idno uint `json:idno` 
	name string `json:name`
	price uint `json:price`
	author Author `json:author`
}	

type Author struct{

	name string `json:name`
	age int `json:age`
	
}


func main(){
   
   auth := Author{name:"javed",age:56}	
   book := Book{idno:123,name:"java",price:450,author:auth}
   
   
   byteArray,err := json.MarshalIndent(book,"","  ")

   if err != nil{
	   fmt.Println(err)
   }

   fmt.Println(string(byteArray[:]))

}