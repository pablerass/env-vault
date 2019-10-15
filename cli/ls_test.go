package cli

import (
    kingpin "gopkg.in/alecthomas/kingpin.v2"

    "github.com/99designs/keyring"
)

func ExampleLsCommand() {
    keyringImpl = keyring.NewArrayKeyring([]keyring.Item{
        {Key: "llamas", Data: []byte(`{"ENV_VAR1":"ABC","ENV_VAR2":"XYZ"}`)},
    })

    app := kingpin.New(`env-vault`, ``)
    ConfigureGlobals(app)
    ConfigureListCommand(app)
    kingpin.MustParse(app.Parse([]string{
        "list",
    }))

    // Output:
    // llamas
}
