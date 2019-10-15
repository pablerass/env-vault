package cli

import (
    kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func ExampleAddCommand() {
    app := kingpin.New(`env-vault`, ``)
    ConfigureGlobals(app)
    ConfigureAddCommand(app)
    kingpin.MustParse(app.Parse([]string{"add", "--debug", "foo"}))

    // Output:
    // Added environment variables to profile "foo" in vault
}
