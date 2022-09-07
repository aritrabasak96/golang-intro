package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}
var mu = sync.RWMutex{} // create a RWMutex instance

func printName() {
	fmt.Println("aritra")
	//mu.RUnlock() // b) once read is complete unlock it
	wg.Done()
}

func printAge() {
	fmt.Println(23)
	//mu.RUnlock() // d) once you complete reading unlock it
	wg.Done()
}

func main() {

	runtime.GOMAXPROCS(1) // 1 thread so the process will go synchronously, do not need to create lock
	for i := 0; i < 12; i++ {

		wg.Add(2)
		//mu.RLock()  // a) lock first
		go printName()
		//mu.RLock() // c) then again lock it because you try to read something
		go printAge()
		wg.Wait()

	}
}

// ------------ NOTE --------------

/**
  Mutex: Mutual exclusion lock
  Sometimes it’s more convenient to synchronize data access by explicit locking.

  Why we use lock => to avoid clashes. (If you have one ticket available and two user come at the same time
  then there will be a clash in it. To avoid this we use lock)

  Mutex{} = in a simple mutex it's just lock or unlock. only one thing can access at a time (lock and unlock)
  RWMutex{} = as many can read at a time but only one can write at a time. If anything is reading then we can't write it at a time

  ... for read lock and unlock -> mu.RLock() & mu.RUnlock()
  ... for write lock and unlock -> mu.Lock() & mu.Unlock()


*/
