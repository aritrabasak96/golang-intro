
package main 

import ("fmt"
        "sync"
        
		)

// create a waitgroup
var wg = sync.WaitGroup{}

// create a channel using make function                                                                                              
var ch1 = make(chan string) // chan => keyword,define that you want to create channel 
                           // string => what type of data you want to pass through the channel  

var ch2 = make(chan string)

func printName()  {
  fmt.Println("aritra")
  ch1 <- "sikdar" // assign (send) data to the channel 
  wg.Done()
}

func printTitle(){
  title := <- ch1 // read data from the channel
  fmt.Println("basak"+title)
  ch2 <- "24" // send data in the channel 
  wg.Done()
}

func printAge(){
  age := <- ch2 // receive data from the channel 
  fmt.Println(age)
  wg.Done()
}
    
func multipleGo(){

    for i:=0;i<10;i++{
          
    wg.Add(3)
    go printAge()
    go printTitle();
    go printName();
    wg.Wait()
    }
}

func main()  {
  
   multipleGo()
}		


// ------------- NOTE -------------

/**
  Most programming language out there are designed to work with single processor 
  But todays computers have multi core processor. Golang is designed to work with multi core processor  
  
  Channels => it is a path or way to communicate and pass data between different goroutines 
   
  // Process, Thread and goroutine 

  processs ==> when you run as many application on your computer at a same time. 
  then you create different processes. Each core of your cpu can handle those processes 
  
  Threads ==> You can create as many threads inside an application. Each core of your cpu handle
  these threads at a time  
 
  // ------------- create channels 

  func receiveOnlyChannel(ch <-chan string){

  }

  func sendOnlyChannel(ch chan<- string){
    
  }

*/