# mx-chain-scenario-cli-go

NewArc's Go CLI for running JSON scenarios through the Go VM backend.

## Runtime boundary

This is deliberately a thin scenario runner. It contains no DRWA, MRV, or
application business rules. Those rules remain in the reviewed NewArc scenario
engine and VM modules, which this repository pins in its root `go.mod`:

- `mx-chain-scenario-go` parses and executes scenarios;
- `mx-chain-vm-go` supplies the current VM host and native-hook integration;
- `mx-chain-vm-v1_4-go` supplies legacy VM compatibility.

The source imports retain the public MultiversX package paths for API
compatibility, while root-level `replace` directives resolve the complete
runtime graph to the corresponding Xorewa revisions. This is necessary because
Go does not inherit `replace` directives from transitive modules.

Do not add a DRWA-specific command-line mode. Chain-coupled DRWA enforcement
must be qualified through the NewArc chain and VM integration tests, where the
native blockchain hook is present.

## Build and verify

Build from a checked-out repository or use a release binary. Direct
`go install module@version` is intentionally not supported: Go ignores or
rejects versioned installs of modules that require local `replace` directives,
and would therefore not preserve this reviewed NewArc dependency graph.

```bash
go mod tidy
go test ./...
go build -ldflags "-X main.buildRevision=local" -o mx-scenario-go ./cmd/mx-scenario-go
./mx-scenario-go version
```

The version output includes the source-build provenance. Release automation
sets `buildRevision` to the Git revision that produced the binary.

## Run a scenario

```bash
./mx-scenario-go run path/to/scenario.scen.json
```

The upstream-compatible flags remain available, including `--vm 1.4`,
`--vm 1.5`, and `--wasmer2`. Use the NewArc chain/VM test suites for scenarios
that require real chain-native behaviour rather than a standalone VM host.
