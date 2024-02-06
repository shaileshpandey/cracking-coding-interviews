package channels

import "fmt"

func SendAMessage(c chan string) {
	c <- "Hello Mr. Pandey"
}

func ReceiveAMessage(c chan string) {
	fmt.Println(<-c)
}
