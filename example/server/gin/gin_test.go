package main

import (
	"testing"

	"github.com/ecoshub/taste"
	example "github.com/ecoshub/taste/example/server"
)

func TestGINServer(t *testing.T) {
	// basic gin server
	s := exampleGINServer()

	// create a tester with server handler and tester
	tester := taste.NewHTTPServerTester(t, s.Handler())

	// run the tester
	tester.Run(example.Scenario)
}
