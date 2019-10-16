package editor

import (
    "strings"

    "github.com/pablerass/env-vault/vault"
)

func GetFromEditor() (vault.EnvVarsSet, error) {
    var envVars = make(vault.EnvVarsSet)
    return EditInEditor(envVars)
}

func EditInEditor(envVars vault.EnvVarsSet) (vault.EnvVarsSet, error) {
    currentEnvVarLines := make([]string, len(envVars))

    line := 0
    for name, value := range envVars {
        currentEnvVarLines[line] = name + "=" + value
        line += 1
    }
    envVarsString, err := EditTextInEditor(strings.Join(currentEnvVarLines, "\n"))
    if err != nil {
        return envVars, err
    }
    updatedEnvVarLines := strings.Split(string(envVarsString), "\n")
    for line = range updatedEnvVarLines {
        if strings.TrimSpace(updatedEnvVarLines[line]) != "" {
            lineSplit := strings.SplitN(updatedEnvVarLines[line], "=", 2)
            envVars[strings.TrimSpace(lineSplit[0])] = strings.TrimSpace(lineSplit[1])
        }
    }
    return envVars, err
}
