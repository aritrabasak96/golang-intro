package main 

import "fmt"

type Write interface{

	Write()
	Write2()
}


type Close interface{
	Close()
}

// this is how you can compose or embed interfaces 
// within an interfaces 
type WriteClose interface{

	Write
	Close 
}

type BufferWriteClose struct{}

func (b1 BufferWriteClose) Write(){
	fmt.Println("b1 executing write")
}

func (b1 BufferWriteClose) Write2(){
	fmt.Println("b1 executing write2")
}

func (b1 BufferWriteClose) Close(){
	fmt.Println("b1 executing close")
}


func main(){

	var w1 WriteClose = BufferWriteClose{}
	w1.Write();
	w1.Close();

	w2  := BufferWriteClose{}
	w2.Write();
	w2.Write2()


}