package server

import (
	"fmt"
	"net/http"
	"testing"
)

type Tester struct {
	t       *testing.T
	handler http.Handler
	ip      string
	store   map[string][]byte
}

func NewTester(t *testing.T, hadler http.Handler) *Tester {
	return &Tester{
		t:       t,
		handler: hadler,
		store:   make(map[string][]byte)}
}

func (tt *Tester) SetIP(ip string) {
	tt.ip = ip
}

func (tt *Tester) Run(scenario []*Case) {
	c, exists := tt.hasOnlyRunMe(scenario)
	if exists {
		fmt.Printf("RUN [ONLY]\t%s\n", c.Name)
		tt.t.Run(c.Name, func(t *testing.T) {
			testTheCase(tt, c, t)
		})
		return
	}

	for _, c := range scenario {
		tt.t.Run(c.Name, func(t *testing.T) {
			testTheCase(tt, c, t)
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
