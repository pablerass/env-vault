package prompt

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

const DefaultEditor = "vim"

func getEditor() string {
	editor := os.Getenv("EDITOR")

	if editor == "" {
		return DefaultEditor
	}

	return editor
}

func knownEditorArguments(executable string, filename string) []string {
	args := []string{filename}

	if strings.Contains(executable, "Visual Studio Code.app") {
		args = append([]string{"--wait"}, args...)
	}

	return args
}

func openFileInEditor(filename string) error {
	executable, err := exec.LookPath(getEditor())
	if err != nil {
		return err
	}

	cmd := exec.Command(executable, knownEditorArguments(executable, filename)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func GetTextFromEditor() (string, error) {
	return EditTextInEditor("")
}

func EditTextInEditor(content string) (string, error) {
	file, err := ioutil.TempFile(os.TempDir(), "*")
	if err != nil {
		return "", err
	}

	filename := file.Name()

	// Defer removal of the temporary file in case any of the next steps fail.
	defer os.Remove(filename)

	if _, err = file.WriteString(content); err != nil {
		return "", err
	}

	if err = file.Close(); err != nil {
		return "", err
	}

	if err = openFileInEditor(filename); err != nil {
		return "", err
	}

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
