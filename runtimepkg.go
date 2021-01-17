package main 

import (
	"fmt"
	"runtime"
)


func defaultCore(){
	
	// runtime.GOMAXPROCS(-1) => by default go returns the number of os threads equals to number of cpu core in your machine 
    // if you have 4 core machine then you have 4 os thread 
	fmt.Printf("Threads----%v",runtime.GOMAXPROCS(-1))
}

func changeCore(){

	// if you want to change the value 
	// this means that you just have 1 thread (no concurrency), totally synchronous 
	runtime.GOMAXPROCS(1)
	fmt.Printf("Threads----%v",runtime.GOMAXPROCS(-1)) // o/p => 1
}


func  main()  {
	
	 defaultCore()
     changeCore()   
}

// -------------------- NOTE ----------------
/**
    runtime.GOMAXPROCS(-1) => by default go returns the number of os threads equals to number of cpu core in your machine 
	if you have 4 core machine then you have 4 os thread 
	
	to check concurrency -> go run -race runtimepkg.go
*/ 
