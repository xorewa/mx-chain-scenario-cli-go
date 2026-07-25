package scencli

import (
	scenclibase "github.com/multiversx/mx-chain-scenario-go/clibase"
)

func ScenariosCLI(version string) {
	scenclibase.ScenariosCLI(version, &runConfig{})
}
