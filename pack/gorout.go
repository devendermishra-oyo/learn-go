package pack

import (
	"fmt"
)

//goroutine is a light weight thread managed by the Go runtime.
// go keyword is used to start a new goroutine.

func say(text string) {
	fmt.Println(text)
}

func inc(num int, c chan int) {
	//c is int channel
	c <- num + 10 //(num+10 is send to channel
}

func LearnGoroutine() {

	//Will start three goroutines.
	//Order is not defined.
	go say("Hello")
	go say("World")
	go say("Good")

	//Channel is used to communicate with goroutines.
	//chan keyword is used to indicate channel
	//channel has two operations: send and receive.
	// <channel> <- val is send operation (channel is on left)
	// val <- <channel> is receive (channel is on right)
	//Also note that the value is copied. Hence, large object should be avoided.

	c := make(chan int) //make is used for multiple operations like
	//creating channel, slices or map.
	go inc(100, c)
	go inc(50, c)

	//Now, we want to read from the channel.
	fmt.Println(<-c, <-c) //Order is not specified.

	//select statements can also be used to wait on multiple channels for send/receive.
}
