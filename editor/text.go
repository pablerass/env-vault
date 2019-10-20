package editor

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

	// Other common editors

	return args
}

func OpenFileInEditor(filename string) error {
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

func GetTextFromEditor() ([]byte, error) {
	return EditTextInEditor("")
}

func EditTextInEditor(content string) ([]byte, error) {
	file, err := ioutil.TempFile(os.TempDir(), "*")
	if err != nil {
		return []byte{}, err
	}

	filename := file.Name()

	// Defer removal of the temporary file in case any of the next steps fail.
	defer os.Remove(filename)

	if _, err = file.WriteString(content); err != nil {
		return []byte{}, err
	}

	if err = file.Close(); err != nil {
		return []byte{}, err
	}

	if err = OpenFileInEditor(filename); err != nil {
		return []byte{}, err
	}

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return []byte{}, err
	}

	return bytes, nil
}
