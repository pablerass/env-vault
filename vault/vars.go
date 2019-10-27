package vault

import (
	"fmt"
	"strings"
)

type EnvVars map[string]string

func NewEnvVars(envVarsString string) (EnvVars, error) {
	var envVars = make(EnvVars)

	envVarLines := strings.Split(envVarsString, "\n")
	for i := range envVarLines {
		if strings.TrimSpace(envVarLines[i]) != "" {
			lineSplit := strings.SplitN(envVarLines[i], "=", 2)
			if len(lineSplit) != 2 {
				return envVars, fmt.Errorf("Invalid format in line %d", i + 1)
			}
			envVars[strings.TrimSpace(lineSplit[0])] = strings.TrimSpace(lineSplit[1])
		}
	}
	return envVars, nil
}

func (envVars EnvVars) String() (string) {
	currentEnvVarLines := make([]string, len(envVars))

	line := 0
	for name, value := range envVars {
		currentEnvVarLines[line] = name + "=" + value
		line += 1
	}
	return strings.Join(currentEnvVarLines, "\n")
}
