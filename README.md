# go-cli

A CLI that filters and logs a specific log level in a log file.

To run this you can simply run the commands `go run .` on the root of this project. This runs the CLI with default flags.

To run with flags, run `go run . -level INFO`, where `-level` is a flag the CLI expects and `INFO` is one of the possible values expected tied to the `-level` flag.

To check the flags and values you can pass to the `go run .` command, run `go run . -help` on the terminal.
