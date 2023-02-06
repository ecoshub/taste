package main

import (
	"testing"

	"github.com/ecoshub/taste"
	example "github.com/ecoshub/taste/example/server"
)

func TestCustomServer(t *testing.T) {
	// basic gin server
	s := exampleDefaultServer()

	// create a tester with server handler and scenario
	tester := taste.NewHTTPServerTester(t, s.Handler())

	// run the scenario
	tester.Run(example.Scenario)
}
