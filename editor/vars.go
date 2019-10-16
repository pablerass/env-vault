package editor

import (
    "fmt"
    "strings"

    "github.com/pablerass/env-vault/vault"
)

func GetFromEditor() (vault.EnvVarsSet, error) {
    var envVars = make(vault.EnvVarsSet)
    return EditInEditor(envVars)
}

func EditInEditor(envVars vault.EnvVarsSet) (vault.EnvVarsSet, error) {
    items := make([]string, len(envVars))

    index := 0
    for name, value := range envVars {
        items[index] = name + "=" + value
        index += 1
        fmt.Println(items)
    }
    varsDefinition, err := EditTextInEditor(strings.Join(items, "\n"))
    if err != nil {
        return envVars, err
    }
    varsDefinitions := strings.Split(string(varsDefinition), "\n")
    for line := range varsDefinitions {
        if strings.TrimSpace(varsDefinitions[line]) != "" {
            lineSplit := strings.SplitN(varsDefinitions[line], "=", 2)
            envVars[strings.TrimSpace(lineSplit[0])] = strings.TrimSpace(lineSplit[1])
        }
    }
    return envVars, err
}
