package cli

import (
	kingpin "gopkg.in/alecthomas/kingpin.v2"

	"github.com/99designs/keyring"
)

func ExampleExecCommand() {
	keyringImpl = keyring.NewArrayKeyring([]keyring.Item{
		{Key: "llamas", Data: []byte(`{"ENV_VAR1":"ABC","ENV_VAR2":"XYZ"}`)},
	})

	app := kingpin.New("env-vault", "")
	ConfigureGlobals(app)
	ConfigureExecCommand(app)
	kingpin.MustParse(app.Parse([]string{
		"--debug", "exec", "llamas", "--", "sh", "-c", "echo $ENV_VAR1",
	}))

	// Output:
	// ABC
}
