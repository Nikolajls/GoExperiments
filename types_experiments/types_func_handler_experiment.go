package types_experiments

import (
	"fmt"
)

// Basic request and response struct
type InRequest struct {
	Payload string
}

type OutResponseWriter struct {
	Load string
}

type ResponseWriter interface {
	Write(toWrite string)
}

func (o *OutResponseWriter) Write(toWrite string) {
	o.Load += toWrite + "\n"
}

type Handler interface {
	GoOn(*InRequest, ResponseWriter)
}

type HandlerFunc func(*InRequest, ResponseWriter)

func (f HandlerFunc) GoOn(in *InRequest, out ResponseWriter) {
	f(in, out)
}

func DoneHandler(a *InRequest, out ResponseWriter) {
	fmt.Println("Running EndHandler")
	out.Write("END")
}
func create() Handler {
	internalHandler := HandlerFunc(DoneHandler)
	return internalHandler
}

// Sample handlers
func AddTest(next Handler) Handler {
	fmt.Println("Creating Test Handler")
	return HandlerFunc(func(in *InRequest, out ResponseWriter) {
		fmt.Println("Running Test Handler")
		out.Write("Test")
		next.GoOn(in, out)
	})
}

// Sample Handler

func AddABC(next Handler) Handler {
	fmt.Println("Creating AddABC Handler")
	return HandlerFunc(func(w *InRequest, out ResponseWriter) {
		fmt.Println("Running AddABC Handler")
		out.Write("ABC")
		next.GoOn(w, out)
	})
}

func AddXXX(next Handler) Handler {
	fmt.Println("Creating AddXXX Handler")
	return HandlerFunc(func(w *InRequest, out ResponseWriter) {
		fmt.Println("Running AddXXX Handler")
		out.Write("XXX")
		next.GoOn(w, out)
	})
}

func HandlePayload(next Handler) Handler {
	fmt.Println("Running HandlePayload Handler")
	return HandlerFunc(func(w *InRequest, out ResponseWriter) {
		out.Write(fmt.Sprintf("Payload: %s", w.Payload))
		next.GoOn(w, out)
	})
}

func CreateInvocation() *InvocationSetup {
	return &InvocationSetup{}
}

func (h *InvocationSetup) AddMiddleware(middlewares ...func(Handler) Handler) *InvocationSetup {
	h.middlewares = append(h.middlewares, middlewares...)
	return h
}

func AddInitialHandler(next Handler) Handler {
	fmt.Println("Creating AddInitialHandler")
	return HandlerFunc(func(w *InRequest, out ResponseWriter) {
		fmt.Println("Running AddInitialHandler")
		out.Write("START")
		next.GoOn(w, out)
	})
}

func EndHandler(a *InRequest, out ResponseWriter) {
	fmt.Println("Running EndHandler")
	out.Write("END")
}

func (h *InvocationSetup) Handle(handler func(Handler) Handler) Invocation {

	LastHandler := HandlerFunc(EndHandler)
	currentMiddleware := handler(LastHandler)

	if len(h.middlewares) > 0 {
		middlewares := h.middlewares
		currentMiddleware = h.middlewares[len(h.middlewares)-1](currentMiddleware)
		for i := len(middlewares) - 2; i >= 0; i-- {
			currentMiddleware = middlewares[i](currentMiddleware)
		}
	}

	handlerChained := AddInitialHandler(currentMiddleware)

	return Invocation{handler: handlerChained}
}

func (h *Invocation) Invoke(input string, rw ResponseWriter) {
	h.handler.GoOn(&InRequest{input}, rw)
}

type InvocationSetup struct {
	middlewares []func(Handler) Handler
	handler     Handler
}

type Invocation struct {
	handler Handler
}
