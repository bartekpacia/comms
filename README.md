# comms

A bunch of command-line utilities to communicate with embedded devices.

> currently it's just 3 files lol

### Usage

Each directory under `cmd/` represents a single binary.

To build a binary, use `go build ./cmd/<name>/*.go`.

To get help, execute the generated binary with the `--help` switch, e.g:

- `./uart_leds --help`
