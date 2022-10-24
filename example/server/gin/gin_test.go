package main

import (
	"testing"

	example "github.com/ecoshub/taste/example/server"
	"github.com/ecoshub/taste/server"
)

func TestGINServer(t *testing.T) {
	// basic gin server
	s := exampleGINServer()

	// create a tester with server handler and scenario
	scenario := server.NewTester(example.Scenario, s.Handler())

	// run the scenario
	scenario.Run(t)
}
