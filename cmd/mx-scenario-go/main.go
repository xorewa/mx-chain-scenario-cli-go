package main

import (
	"fmt"

	scencli "github.com/xorewa/mx-chain-scenario-cli-go/cli"
)

const version = "5.1.0-xorewa"

// buildRevision is set by release builds with:
// -ldflags "-X main.buildRevision=<source-revision>".
// Keeping the default explicit prevents a locally built binary from claiming a
// release revision it does not contain.
var buildRevision = "dev"

func main() {
	scencli.ScenariosCLI(fmt.Sprintf("%s (build %s)", version, buildRevision))
}
