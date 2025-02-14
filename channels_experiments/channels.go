package channels_experiments

import (
	"fmt"
	"time"
)

func DoChannelTests() {
	TestChannelDirections()
	//TestStructChannel()
}
func SendMessage[T any](aSendOnlyChannel chan<- T, message T) {
	aSendOnlyChannel <- message
	fmt.Printf("Sending message on channel %v\n", message)
}

func ReceiveMessage[T any](aReceiveOnlyChannel <-chan T) T {
	receivedMessage := <-aReceiveOnlyChannel
	fmt.Printf("Received message on channel %v\n", receivedMessage)
	return receivedMessage
}

func ForwardMessage[T any](receiveChan <-chan T, forward chan<- T) {
	message := ReceiveMessage(receiveChan)
	fmt.Printf("Forwarding message:%v on channel\n", message)
	SendMessage(forward, message)
}

func TestChannelDirections() {
	fmt.Println("TestChannelDirections start")
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	SendMessage(pings, "A lovely Message to send")
	ForwardMessage(pings, pongs)
	pongReceived := <-pongs
	fmt.Printf("Received message on channel %v\n", pongReceived)

}

func SendMessageToChannelWithNoReceiver(bufferingAmount int, message string) <-chan string {
	var ch chan string
	if bufferingAmount > 0 {
		ch = make(chan string, bufferingAmount)
	} else {
		ch = make(chan string)
	}

	fmt.Printf("%v\tSending message on channel\n", time.Now())
	ch <- message
	fmt.Printf("%v\tSENT message on channel\n", time.Now())

	msgCount := len(ch)
	fmt.Printf("%v\tSENT message on channel\n", msgCount)
	//time.Sleep(time.Second * 2)
	return ch
}
