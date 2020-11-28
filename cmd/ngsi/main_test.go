package main

import (
	"flag"
	"testing"

	"github.com/google/go-cmdtest"
)

var update = flag.Bool("help", false, "update test files with results")

func TestHelp(t *testing.T) {
	ts, err := cmdtest.Read("testdata/help")
	if err != nil {
		t.Fatal(err)
	}
	ts.Commands["ngsi"] = cmdtest.InProcessProgram("ngsi", run)
	ts.Run(t, *update)
}
