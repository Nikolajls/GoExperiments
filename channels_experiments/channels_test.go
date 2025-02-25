package channels_experiments

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_SendMessageToChannelWithNoReceiver_Buffer(t *testing.T) {
	ch := SendMessageToChannelWithNoReceiver(1, "TEST")
	//msg := <-ch
	msgCount := len(ch)
	assert.Equal(t, 1, msgCount)

}

func Test_SendingOnChannelWithNoReceiver_WillBlock(t *testing.T) {
	lotteryNumbers := make(chan int)
	doneDrawing := make(chan struct{})

	go func() {
		lotteryNumbers <- 42 // This will block because there is no receiver for the channel
		close(doneDrawing)   // Since the above is blocking in this go subroutine meaning the closing of doneDrawing channel is never executed.
	}()

	select {
	case <-doneDrawing:
		t.Error("Expected the send operation to block, but it did not.")
	case <-time.After(1 * time.Second):
		t.Log("Send operation is blocking as expected.")
	}
}

func Test_SendingOnChannelWithReceiver_WillNotBlock(t *testing.T) {
	lotteryNumbers := make(chan int)
	doneDrawing := make(chan struct{})

	go func() {
		lotteryNumbers <- 42 // Contrary to the test above we configure a channel consumer listener below meaning that this will not block
		close(doneDrawing)   // Since the above is no longer blocking we can call close on secondary doneDrawing channel now which is used in the select statement
	}()

	go func() {
		lotteryNumber := <-lotteryNumbers
		fmt.Println("Received from lottery channel", lotteryNumber)
	}()

	select {
	case <-doneDrawing:
		t.Log("Done drawing numbers")
	case <-time.After(1 * time.Second):
		// If we reach here, it means the send operation is blocking as expected
		t.Error("Send operation is blocking but is not what is expected.")
	}
}
