package main

import (
	"testing"

	"github.com/ecoshub/taste"
	"github.com/ecoshub/taste/example"
)

func TestCustomServer(t *testing.T) {
	// basic gin server
	server := exampleDefaultServer()

	// create a tester with server handler and scenario
	tester := taste.NewTester(server.Handler(), example.Scenario)

	// run the scenario
	tester.Run(t)
}
