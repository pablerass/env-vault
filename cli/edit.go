package cli

import (
	"fmt"

	"github.com/99designs/keyring"
	"github.com/pablerass/env-vault/prompt"
	"github.com/pablerass/env-vault/vault"
	"gopkg.in/alecthomas/kingpin.v2"
)

type EditCommandInput struct {
	Profile string
	Keyring keyring.Keyring
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

	text, err := prompt.EditTextInEditor(envVars.String())
	if err != nil {
		app.Fatalf(err.Error())
		return
	}

	envVars, err = vault.NewEnvVars(text)
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
