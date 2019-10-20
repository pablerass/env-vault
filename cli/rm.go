package cli

import (
	"fmt"

	"github.com/99designs/keyring"
	"github.com/pablerass/env-vault/prompt"
	"github.com/pablerass/env-vault/vault"
	"gopkg.in/alecthomas/kingpin.v2"
)

type RemoveCommandInput struct {
	Profile string
	Keyring keyring.Keyring
}

func ConfigureRemoveCommand(app *kingpin.Application) {
	input := RemoveCommandInput{}

	cmd := app.Command("remove", "Removes environment variables")
	cmd.Alias("rm")

	cmd.Arg("profile", "Name of the profile").
		Required().
		StringVar(&input.Profile)

	cmd.Action(func(c *kingpin.ParseContext) error {
		input.Keyring = keyringImpl
		RemoveCommand(app, input)
		return nil
	})
}

func RemoveCommand(app *kingpin.Application, input RemoveCommandInput) {
	provider := &vault.KeyringProvider{Keyring: input.Keyring, Profile: input.Profile}
	r, err := prompt.TerminalPrompt(fmt.Sprintf("Delete environment variables for profile %q? (Y|n)", input.Profile))
	if err != nil {
		app.Fatalf(err.Error())
		return
	} else if r == "N" || r == "n" {
		return
	}

	if err := provider.Delete(); err != nil {
		app.Fatalf(err.Error())
		return
	}
	fmt.Printf("Deleted profile.\n")
}
