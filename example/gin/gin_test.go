package main

import (
	"testing"

	"github.com/ecoshub/taste"
	"github.com/ecoshub/taste/example"
)

func TestGINServer(t *testing.T) {
	// basic gin server
	server := exampleGINServer()

	// create a default scenario
	scenario := taste.NewScenario()

	// add test cases to scenario
	scenario.AddCases(example.TestCases)

	// attach server to scenario
	scenario.AttachCustomServer(server.Handler())

	// run the scenario
	scenario.Run(t)
}
