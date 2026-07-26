package main

import (
	"fmt"

	scencli "github.com/xorewa/mx-chain-scenario-cli-go/cli"
)

// Keep this in sync with the release tag (without the leading `v`).  The
// installer pins a release tag, so reporting a different binary version makes
// incident diagnosis and provenance checks needlessly ambiguous.
const version = "5.1.2-xorewa.1"

// buildRevision is set by release builds with:
// -ldflags "-X main.buildRevision=<source-revision>".
// Keeping the default explicit prevents a locally built binary from claiming a
// release revision it does not contain.
var buildRevision = "dev"

func main() {
	scencli.ScenariosCLI(fmt.Sprintf("%s (build %s)", version, buildRevision))
}
