package taste

import (
	"fmt"
	"net/http"
	"testing"
)

type Scenario struct {
	server    http.Handler
	mockIP    string
	runBefore func(t *testing.T)
	runAfter  func(t *testing.T)
	cases     []*Case
}

func NewScenario() *Scenario {
	return &Scenario{cases: make([]*Case, 0, 8)}
}

func (sc *Scenario) AddCases(c []*Case) {
	sc.cases = c
}

func (sc *Scenario) AddCase(c *Case) {
	sc.cases = append(sc.cases, c)
}

func (sc *Scenario) AttachServer(server *HTTPMockServer) {
	sc.server = server.mux
}

func (sc *Scenario) AttachCustomServer(server http.Handler, mockIPoptional ...string) {
	mockIP := defaultMockIP
	if len(mockIPoptional) == 1 {
		mockIP = mockIPoptional[0]
	}
	sc.mockIP = mockIP
	sc.server = server
}

func (sc *Scenario) Run(t *testing.T) {
	c, exists := sc.hasOnlyRunMe()
	if exists {
		fmt.Printf("RUN [ONLY]\t%s\n", c.Name)
		t.Run(c.Name, func(t *testing.T) {
			run(sc, c, t)
		})
		return
	}

	for _, c := range sc.cases {
		t.Run(c.Name, func(t *testing.T) {
			run(sc, c, t)
		})
	}
}
