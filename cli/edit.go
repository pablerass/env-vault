package cli

import (
    "fmt"
    "strings"

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

    items := make([]string, len(envVars))

    index := 0
    for name, value := range envVars {
        items[index] = name + "=" + value
        index += 1
    }
    varsDefinition, _ := editor.EditInEditor(strings.Join(items, "\n"))
    varsDefinitions := strings.Split(string(varsDefinition), "\n")
    for line := range varsDefinitions {
        if strings.TrimSpace(varsDefinitions[line]) != "" {
            lineSplit := strings.SplitN(varsDefinitions[line], "=", 2)
            fmt.Println(lineSplit)
            envVars[strings.TrimSpace(lineSplit[0])] = strings.TrimSpace(lineSplit[1])
        }
    }
    if err := provider.Store(envVars); err != nil {
        app.Fatalf(err.Error())
        return
    }

    fmt.Printf("Edited environment variables to profile %q in vault\n", input.Profile)
}
