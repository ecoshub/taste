package main

import (
	"testing"

	example "github.com/ecoshub/taste/example/server"
	"github.com/ecoshub/taste/server"
)

func TestCustomServer(t *testing.T) {
	// basic gin server
	s := exampleDefaultServer()

	// create a tester with server handler and scenario
	tester := server.NewTester(example.Scenario, s.Handler())

	// run the scenario
	tester.Run(t)
}
