package vault

import (
	"fmt"
)

func ExampleNewEnvVars() {
	var err error

	envVars, _ := NewEnvVars("")
	fmt.Println(len(envVars))

	envVars, _ = NewEnvVars("A=A")
	fmt.Println(len(envVars))
	fmt.Println(envVars["A"])

	envVars, _ = NewEnvVars("A=A\nB=B\n")
	fmt.Println(len(envVars))
	fmt.Println(envVars["B"])
	fmt.Println(envVars["A"])

	_, err = NewEnvVars("A=A\nBB\n")
	fmt.Println(err)

	// Output:
	// 0
	// 1
	// A
	// 2
	// B
	// A
	// Invalid format in line 2
}
