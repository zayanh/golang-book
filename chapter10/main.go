package main
import (
	"fmt"
	"time"
	"math/rand"
)

func f(n int) {
	for i := 0; i < 4; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func pinger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func printer(c <-chan string) {
	for {
		fmt.Println(<-c)
		time.Sleep(time.Second * 1)
	}
}

func manualSleep(i int) {
	// Is there a better way to make a blocking call to a channel?
	select { case <- time.After(time.Second * time.Duration(i)):}
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	fmt.Println("press any key to exit")

	//for i:= 0; i < 4; i++ {
	//	go f(i)
	//}

	//var c chan string = make(chan string)

	//go pinger(c)
	//go ponger(c)
	//go printer(c)

	go func() {
		for {
			c1 <- "from 1"
			//time.Sleep(time.Second * 2)
			manualSleep(2)
		}
	}()

	go func() {
		for {
			c1 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		for {
			// select is a blocking call - it will wait for one of
			//   the channels to be ready if none are. If multiple
			//   are ready one will be picked randomly
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			case <- time.After(time.Second):
				fmt.Println("channels timed out")
				return
			//default:
			//	fmt.Println("nothing ready")
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}
