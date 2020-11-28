package main

import (
	"flag"
	"testing"

	"github.com/google/go-cmdtest"
	"github.com/lets-fiware/ngsi-go/internal/ngsicmd"
)

var update = flag.Bool("help", false, "update test files with results")

func TestHelp(t *testing.T) {
	ts, err := cmdtest.Read("testdata/help")
	if err != nil {
		t.Fatal(err)
	}
	ts.Commands["ngsi"] = cmdtest.InProcessProgram("ngsi", ngsicmd.Run)
	ts.Run(t, *update)
}
