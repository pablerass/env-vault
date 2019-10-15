package main

import (
    "os"

    "github.com/pablerass/env-vault/cli"
    "gopkg.in/alecthomas/kingpin.v2"
)

// Version is provided at compile time
var Version = "dev"

func main() {
    run(os.Args[1:], os.Exit)
}

func run(args []string, exit func(int)) {
    app := kingpin.New(
        `env-vault`,
        `A vault for securely storing and accessing environment variable sets.`,
    )

    app.Writer(os.Stdout)
    app.Version(Version)
    app.Terminate(exit)

    cli.ConfigureGlobals(app)
    cli.ConfigureAddCommand(app)
    cli.ConfigureListCommand(app)
    cli.ConfigureExecCommand(app)
    cli.ConfigureRemoveCommand(app)

    kingpin.MustParse(app.Parse(args))
}