package server

import (
	"fmt"
	"net/http"
	"testing"
)

type Tester struct {
	scenario scenario

	handler http.Handler
	ip      string
	store   map[string][]byte
}

func NewTester(sc scenario, optionalHandler ...http.Handler) *Tester {
	var h http.Handler
	if len(optionalHandler) == 0 {
		h = nil
	} else {
		h = optionalHandler[0]
	}
	return &Tester{
		scenario: sc,
		handler:  h,
		store:    make(map[string][]byte)}
}

func (tt *Tester) AttachHandler(handler http.Handler) {
	tt.handler = handler
}

func (tt *Tester) Run(t *testing.T) {
	if tt.handler == nil {
		t.Fatal("there is  no handler to test this scenario. please attach a handler with 'AttachHandler' function")
	}
	c, exists := tt.hasOnlyRunMe()
	if exists {
		fmt.Printf("RUN [ONLY]\t%s\n", c.Name)
		t.Run(c.Name, func(t *testing.T) {
			run(tt, c, t)
		})
		return
	}

	for _, c := range tt.scenario {
		t.Run(c.Name, func(t *testing.T) {
			run(tt, c, t)
		})
	}
}

func (tt *Tester) ResetStore() {
	tt.store = make(map[string][]byte)
}

func (tt *Tester) Store(key string, body []byte) {
	tt.store[key] = body
}

func (tt *Tester) StoreKeyValue(key, value string) {
	tt.Store(key, []byte(value))
}
