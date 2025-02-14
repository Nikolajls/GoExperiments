package channels_experiments

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SendMessageToChannelWithNoReceiver_Buffer(t *testing.T) {
	ch := SendMessageToChannelWithNoReceiver(1, "TEST")
	//msg := <-ch
	msgCount := len(ch)
	assert.Equal(t, 1, msgCount)

}
func Test_SendMessageToChannelWithNoReceiver_NoBuffer(t *testing.T) {
	ch := SendMessageToChannelWithNoReceiver(0, "TEST")
	msg := <-ch
	msgCount := len(ch)
	assert.Equal(t, 1, msgCount)
	assert.Equal(t, "TEST", msg)

}
