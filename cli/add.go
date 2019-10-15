package cli

import (
    "fmt"

    "github.com/pablerass/env-vault/vault"
    "github.com/99designs/keyring"
    "gopkg.in/alecthomas/kingpin.v2"
)

type AddCommandInput struct {
    Profile   string
    Keyring   keyring.Keyring
}

func ConfigureAddCommand(app *kingpin.Application) {
    input := AddCommandInput{}

    cmd := app.Command("add", "Adds environment variables profile")
    cmd.Arg("profile", "Name of the profile").
        Required().
        StringVar(&input.Profile)

    cmd.Action(func(c *kingpin.ParseContext) error {
        input.Keyring = keyringImpl
        AddCommand(app, input)
        return nil
    })
}

func AddCommand(app *kingpin.Application, input AddCommandInput) {
    envVars := "value"
    provider := &vault.KeyringProvider{Keyring: input.Keyring, Profile: input.Profile}

    if err := provider.Store(envVars); err != nil {
        app.Fatalf(err.Error())
        return
    }

    fmt.Printf("Added environment variables to profile %q in vault\n", input.Profile)
}
