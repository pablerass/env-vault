package cli

import (
    "fmt"
    "os"
    "text/tabwriter"

    "github.com/99designs/keyring"
    "gopkg.in/alecthomas/kingpin.v2"
)

type LsCommandInput struct {
    Keyring         keyring.Keyring
}

func ConfigureListCommand(app *kingpin.Application) {
    input := LsCommandInput{}

    cmd := app.Command("list", "List profiles")
    cmd.Alias("ls")

    cmd.Action(func(c *kingpin.ParseContext) error {
        input.Keyring = keyringImpl
        LsCommand(app, input)
        return nil
    })
}

func LsCommand(app *kingpin.Application, input LsCommandInput) {
    profiles, err := input.Keyring.Keys()
    if err != nil {
        app.Fatalf(err.Error())
        return
    }

    if len(profiles) == 0 {
        app.Fatalf("No profiles found")
        return
    }

    w := tabwriter.NewWriter(os.Stdout, 25, 4, 2, ' ', 0)
    fmt.Fprintln(w, "Profile")
    fmt.Fprintln(w, "=======")

    for _, profile := range profiles {
        fmt.Fprintf(w, "%s", profile)
    }

    if err = w.Flush(); err != nil {
        app.Fatalf("%v", err)
        return
    }
}
