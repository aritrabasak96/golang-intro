package main 

import ("fmt"
        "sync")

var wg = sync.WaitGroup{}
// create a buffered channel 
var ch = make(chan int, 1) 
var ch2 = make(chan int,3)


func oldReceiver(ch <-chan int){
	
	// receive the two value 
	value := <- ch // receive the first value from sender 
	fmt.Println("value----",value)
	value = <- ch // receive the second value fron the sender 
	fmt.Println("value 2-----",value)
	value = <- ch // receive the third value fron the sender 
	fmt.Println("value 3-----",value)
	wg.Done()
}

// this is the best way to receive multiple values from channel 
func newReceiver(ch <-chan int){

	for i:= range ch{
		fmt.Println(i)
	}
	wg.Done()
}

func sender(ch chan<- int){

	// send 3 integer value one by one 
	ch <- 12;
	ch <- 15;
	ch <- 20
	close(ch); // you explicietely tell, that you are done with adding data into the channel
	wg.Done()
}

func multipleChannel(){
	
	wg.Add(2)
	go sender(ch)
	go oldReceiver(ch)
	wg.Wait()
}

func main(){
	
	multipleChannel();
   
	// another example 
	// you fill all the buffer size of the channel ch2
	ch2 <- 12;
	ch2 <- 18;
	ch2 <- 22;

	fmt.Println(<-ch2)
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)
}