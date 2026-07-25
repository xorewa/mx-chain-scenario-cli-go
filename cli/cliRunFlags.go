package scencli

import (
	"fmt"

	scenclibase "github.com/multiversx/mx-chain-scenario-go/clibase"
	scenexec "github.com/multiversx/mx-chain-scenario-go/scenario/executor"
	scenio "github.com/multiversx/mx-chain-scenario-go/scenario/io"
	vm15scenario "github.com/multiversx/mx-chain-vm-go/scenario"
	vm15wasmer2 "github.com/multiversx/mx-chain-vm-go/wasmer2"
	vm14scenario "github.com/multiversx/mx-chain-vm-v1_4-go/scenario"
	cli "github.com/urfave/cli/v2"
)

const vmFlag = "vm"
const vm14FlagValue = "1.4"
const vm15FlagValue = "1.5"
const vmDefaultFlagValue = vm15FlagValue

var _ scenclibase.CLIRunConfig = (*runConfig)(nil)

type runConfig struct{}

func (*runConfig) GetFlags() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:    "force-trace-gas",
			Aliases: []string{"g"},
			Usage:   "overrides the traceGas option in the scenarios`",
		},
		&cli.StringFlag{
			Name:  "vm",
			Usage: "allows to select the VM to run (1.4 | 1.5)`",
		},
		&cli.BoolFlag{
			Name:  "wasmer1",
			Usage: "use the wasmer1 executor (no longer available)`",
		},
		&cli.BoolFlag{
			Name:  "wasmer2",
			Usage: "use the wasmer2 executor`",
		},
	}
}

func parseVMFlag(cCtx *cli.Context) string {
	vmFlagStr := cCtx.String(vmFlag)
	switch vmFlagStr {
	case "":
		return vmDefaultFlagValue
	case vm15FlagValue:
		return vm15FlagValue
	case vm14FlagValue:
		return vm14FlagValue
	default:
		panic(fmt.Sprintf("invalid vm flag: %s", vmFlagStr))
	}
}

func (*runConfig) ParseFlags(cCtx *cli.Context) scenclibase.CLIRunOptions {
	runOptions := &scenio.RunScenarioOptions{
		ForceTraceGas: cCtx.Bool("force-trace-gas"),
	}

	var vmBuilder scenexec.VMBuilder
	vmFlagStr := parseVMFlag(cCtx)
	switch vmFlagStr {
	case vm15FlagValue:
		vm15Builder := vm15scenario.NewScenarioVMHostBuilder()
		if cCtx.Bool("wasmer1") {
			panic("wasmer1 no longer available")
		}
		if cCtx.Bool("wasmer2") {
			vm15Builder.OverrideVMExecutor = vm15wasmer2.ExecutorFactory()
		}
		vmBuilder = vm15Builder
	case vm14FlagValue:
		vmBuilder = vm14scenario.NewScenarioVMHostBuilder()
	default:
		panic(fmt.Sprintf("invalid vm flag: %s", vmFlagStr))
	}

	return scenclibase.CLIRunOptions{
		RunOptions: runOptions,
		VMBuilder:  vmBuilder,
	}
}
