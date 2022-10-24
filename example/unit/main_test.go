package main

import (
	"testing"

	"github.com/ecoshub/taste/server"
	"github.com/ecoshub/taste/unit"
)

func TestNewTester(t *testing.T) {
	unit.NewCase().Name("server.NewTester").Run(server.NewTester([]*server.Case{})).Expect((*server.Tester)(nil)).Test(t)
}
