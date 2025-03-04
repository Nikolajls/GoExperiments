package types_experiments

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockedResponseWriter struct {
	writes []string
}

func (o *MockedResponseWriter) Write(toWrite string) {
	o.writes = append(o.writes, toWrite)
}

func TestInvocation_Invoke_WithMiddlewares(t *testing.T) {
	invocationSetup := CreateInvocation().AddMiddleware(AddTest)
	invocation := invocationSetup.Handle(HandlePayload)
	r := &MockedResponseWriter{}
	invocation.Invoke("PayloadGoesHere", r)
	fmt.Println(r.writes)
	t.Log(r.writes)
	assert.EqualValues(t, []string{"START", "Test", "Payload: PayloadGoesHere", "END"}, r.writes)
}

func TestInvocation_Invoke_WithoutMiddlewares(t *testing.T) {
	invocationSetup := CreateInvocation()
	invocation := invocationSetup.Handle(HandlePayload)
	r := &MockedResponseWriter{}
	invocation.Invoke("PayloadGoesHere", r)
	fmt.Println(r.writes)
	t.Log(r.writes)
	assert.EqualValues(t, []string{"START", "Payload: PayloadGoesHere", "END"}, r.writes)
}
