package cli

import (
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"

	"github.com/99designs/keyring"
	"github.com/pablerass/env-vault/vault"
	"gopkg.in/alecthomas/kingpin.v2"
)

type ExecCommandInput struct {
	Profile string
	Command string
	Args    []string
	Keyring keyring.Keyring
	Signals chan os.Signal
}

func ConfigureExecCommand(app *kingpin.Application) {
	input := ExecCommandInput{}

	cmd := app.Command("exec", "Executes a command with AWS credentials in the environment")

	cmd.Arg("profile", "Name of the profile").
		Required().
		StringVar(&input.Profile)

	cmd.Arg("cmd", "Command to execute").
		Default(os.Getenv("SHELL")).
		StringVar(&input.Command)

	cmd.Arg("args", "Command arguments").
		StringsVar(&input.Args)

	cmd.Action(func(c *kingpin.ParseContext) error {
		input.Keyring = keyringImpl
		input.Signals = make(chan os.Signal)
		ExecCommand(app, input)
		return nil
	})
}

func ExecCommand(app *kingpin.Application, input ExecCommandInput) {
	if os.Getenv("ENV_VAULT") != "" {
		app.Fatalf("env-vault sessions should be nested with care, unset $ENV_VAULT to force")
		return
	}

	// TODO: Check if profile exists

	provider := &vault.KeyringProvider{Keyring: input.Keyring, Profile: input.Profile}

	envVars, err := provider.Retrieve()
	if err != nil {
		app.Fatalf(err.Error())
		return
	}

	env := environ(os.Environ())
	env.Set("ENV_VAULT", input.Profile)

	for name, value := range envVars {
		env.Set(name, value)
	}
	cmd := exec.Command(input.Command, input.Args...)
	cmd.Env = env
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	signal.Notify(input.Signals, os.Interrupt, os.Kill)

	if err := cmd.Start(); err != nil {
		app.Fatalf("%v", err)
	}
	// wait for the command to finish
	waitCh := make(chan error, 1)
	go func() {
		waitCh <- cmd.Wait()
		close(waitCh)
	}()

	for {
		select {
		case sig := <-input.Signals:
			if err = cmd.Process.Signal(sig); err != nil {
				app.Errorf("%v", err)
				break
			}
		case err := <-waitCh:
			var waitStatus syscall.WaitStatus
			if exitError, ok := err.(*exec.ExitError); ok {
				waitStatus = exitError.Sys().(syscall.WaitStatus)
				os.Exit(waitStatus.ExitStatus())
			}
			if err != nil {
				app.Fatalf("%v", err)
			}
			return
		}
	}
}

// environ is a slice of strings representing the environment, in the form "key=value".
type environ []string

// Unset an environment variable by key
func (e *environ) Unset(key string) {
	for i := range *e {
		if strings.HasPrefix((*e)[i], key+"=") {
			(*e)[i] = (*e)[len(*e)-1]
			*e = (*e)[:len(*e)-1]
			break
		}
	}
}

// Set adds an environment variable, replacing any existing ones of the same key
func (e *environ) Set(key, val string) {
	e.Unset(key)
	*e = append(*e, key+"="+val)
}
