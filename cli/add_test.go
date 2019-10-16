package cli

import (
    "os"

    kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func ExampleAddCommand() {
	os.Setenv("ENV_VAULT_BACKEND", "file")
	os.Setenv("ENV_VAULT_FILE_PASSPHRASE", "password")

	defer os.Unsetenv("ENV_VAULT_BACKEND")
	defer os.Unsetenv("ENV_VAULT_FILE_PASSPHRASE")

    app := kingpin.New(`env-vault`, ``)
    ConfigureGlobals(app)
    ConfigureAddCommand(app)
    kingpin.MustParse(app.Parse([]string{"add", "--debug", "foo"}))

    // Output:
    // Added environment variables to profile "foo" in vault
}
