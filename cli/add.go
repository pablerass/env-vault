package cli

import (
    "fmt"
    "strings"

    "github.com/pablerass/env-vault/editor"
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
    envVars := make(vault.EnvVarsSet)
    varsDefinition, _ := editor.GetFromEditor()
    varsDefinitions := strings.Split(string(varsDefinition), "\n")
    for line := range varsDefinitions {
        if strings.TrimSpace(varsDefinitions[line]) != "" {
            lineSplit := strings.SplitN(varsDefinitions[line], "=", 2)
            fmt.Println(lineSplit)
            envVars[strings.TrimSpace(lineSplit[0])] = strings.TrimSpace(lineSplit[1])
        }
    }
    provider := &vault.KeyringProvider{Keyring: input.Keyring, Profile: input.Profile}

    if err := provider.Store(envVars); err != nil {
        app.Fatalf(err.Error())
        return
    }

    fmt.Printf("Added environment variables to profile %q in vault\n", input.Profile)
}
