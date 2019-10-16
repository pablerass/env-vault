package cli

import (
    "fmt"

    "github.com/pablerass/env-vault/editor"
    "github.com/pablerass/env-vault/vault"
    "github.com/99designs/keyring"
    "gopkg.in/alecthomas/kingpin.v2"
)

type EditCommandInput struct {
    Profile   string
    Keyring   keyring.Keyring
}

func ConfigureEditCommand(app *kingpin.Application) {
    input := EditCommandInput{}

    cmd := app.Command("edit", "Edit profile environment variables")
    cmd.Arg("profile", "Name of the profile").
        Required().
        StringVar(&input.Profile)

    cmd.Action(func(c *kingpin.ParseContext) error {
        input.Keyring = keyringImpl
        EditCommand(app, input)
        return nil
    })
}

func EditCommand(app *kingpin.Application, input EditCommandInput) {
    provider := &vault.KeyringProvider{Keyring: input.Keyring, Profile: input.Profile}
    envVars, err := provider.Retrieve()
    if err != nil {
        app.Fatalf(err.Error())
        return
    }

    envVars, err = editor.EditInEditor(envVars)
    if err != nil {
        app.Fatalf(err.Error())
        return
    }
    if err := provider.Store(envVars); err != nil {
        app.Fatalf(err.Error())
        return
    }

    fmt.Printf("Edited environment variables to profile %q in vault\n", input.Profile)
}
