package taste

import (
	"fmt"
	"net/http"
	"testing"
)

type Tester struct {
	Scenario scenario

	handler http.Handler
	ip      string
}

func NewTester(handler http.Handler, sc scenario) *Tester {
	return &Tester{Scenario: sc, handler: handler}
}

func (tt *Tester) Run(t *testing.T) {
	c, exists := tt.hasOnlyRunMe()
	if exists {
		fmt.Printf("RUN [ONLY]\t%s\n", c.Name)
		t.Run(c.Name, func(t *testing.T) {
			run(tt, c, t)
		})
		return
	}

	for _, c := range tt.Scenario {
		t.Run(c.Name, func(t *testing.T) {
			run(tt, c, t)
		})
	}
}
