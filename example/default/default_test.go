package main

import (
	"testing"

	"github.com/ecoshub/taste"
	"github.com/ecoshub/taste/example"
)

func TestCustomServer(t *testing.T) {
	// basic gin server
	server := exampleDefaultServer()

	// create a default scenario
	scenario := taste.NewScenario()

	// add test cases to scenario
	scenario.AddCases(example.TestCases)

	// attach server to scenario
	scenario.AttachServer(server)

	// run the scenario
	scenario.Run(t)
}
