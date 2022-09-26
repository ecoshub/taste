package main

import (
	"testing"

	"github.com/ecoshub/taste"
	"github.com/ecoshub/taste/example"
)

func TestGINServer(t *testing.T) {
	// basic gin server
	server := exampleGINServer()

	// create a tester with server handler and scenario
	scenario := taste.NewTester(example.Scenario, server.Handler())

	// run the scenario
	scenario.Run(t)
}
