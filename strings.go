package main 

import "fmt"

func strings(){
	
	   n := "hello aritra 😍😍 ok😊😊"
	   t := []byte(n)   // convert the string to array's of byte 
       rune_ := []rune(n)

	   fmt.Printf("%v,%T",t,t)  // uint8 => unsigned integer range between 0 to 255 
	   fmt.Printf("%v,%T",rune_,rune_)
}


func main(){

   strings();
}

// ------------------ NOTES --------------------

// byte is a alias for unit8
// rune is a alias for int32
// every character is a rune ex: 'a', '😍'
// anything which can be encoded as 4 byte is a rune