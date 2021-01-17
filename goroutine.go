/**
 what is the difference between parallelism and concurrency

 1) parallelism  => Doing multiple task at a same time and each tast is independent
 with each other. (Thread) (You should have multicore processor)
 
 2) concurrency => Dealing multiple things at once (works for single core processor)


*/

package main 

import ("fmt" 
	    "sync")


var wg = sync.WaitGroup{};



func goParent1(){

	msg := "go parent 1 start";
	
	go func(msg2 string){

		fmt.Println("child----",msg2)
		// once you finish the execution 
		wg.Done()
	}(msg)
	 
	
	msg = "go Parent 1 end"
   
}


func goParent2(){

	msg := "go parent 2 start";
	
	go func(msg2 string){

		fmt.Println("child----",msg2)
		// once you finish the execution 
		wg.Done()
	}(msg)
	 
	
	msg = "go Parent 2 end";
   
}

func main(){
	
	wg.Add(2); // we have two go routine to execute 
	goParent1();  // you can not predict the output 
	goParent2();
	wg.Wait()
}

// -------------------- NOTE -------------------
/**
  1) goroutine is a light weight, abstraction of a thread
  1.1) You can create as many goroutines inside a program  
  2) when you run the main function => the main function is actually executing in a goroutine itself 
  so when the main function finish it's execution, it does not care about other goroutine
  
  3) time.Sleep(1000 * time.Millisecond) is not a good practice in production level 
     instead you can create waitGroup
*/

