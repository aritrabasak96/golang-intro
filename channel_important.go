/*
   ~ what is all goroutines are asleep - deadlock error

   ~ when should we use buffered channel

   ~ NOTE: This is very important that you should close the channel before main function exits

        A common way of doing is that =

		defer func(){
			close(ch); // close the channel
		}()

*/

package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func causeDeadlock() {

	/*
		    ~ here you can see that we have only one channel and we can pass only string type data

			~ send send a message (1)

			~ receiver receive the message and print it and finish its execution

			~ after 10 seconds sender send the message again, but this time no receiver is there to process (2)

			~ that is the reason of this error: what is all goroutines are asleep - deadlock error

			~ (2) causes the deadlock because we append message to the channel but no one to process it, so the whole
			 execution is blocked there - go runtime understand that and throws deadlock error
	*/

	//  condition 1
	ch := make(chan string)

	wg.Add(2)
	// receiver
	go func() {

		value := <-ch
		fmt.Println("first value from channel- " + value)
		wg.Done()

	}()

	// sender
	go func() {

		ch <- "hello" // (1)
		fmt.Println("sender send value")
		time.Sleep(10 * time.Second)
		ch <- "bye" // (2)
		fmt.Println("send send value again")
		wg.Done()
	}()

}

func bufferedChannel() {

	/*

		    ~ Now sender send messages frequently and receiver needs some time to process those messages
			- That is the ideal scenario where you should use buffered channel

			~ (1) => close(ch) = if you do not close this channel then (2) will throw a deadlock error.
			why? = Because for range loop will expect to have more data to process, but it could not find it

			~ (1) = so we need to close the channel and (2) for-range will detect that no more data to process so
			it will exit

	*/

	ch := make(chan int, 30) // we create a buffered channel, 30 = we can store up to 30 integers

	// create a receiver
	go func(ch <-chan int) {

		for i := range ch {
			// (2)
			time.Sleep(10 * time.Second) // some time to process
			fmt.Println("your processed message is" + strconv.Itoa(i))
		}
		wg.Done()
	}(ch)

	// sender
	go func(ch chan<- int) {

		ch <- 10
		ch <- 12
		close(ch) // (1)
		wg.Done()
	}(ch)
}

func main() {

	causeDeadlock()
	wg.Wait()
}
