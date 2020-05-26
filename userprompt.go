package userprompt

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"io"
	"os"
	"strings"
)

// So tests can override
var (
	Input  io.Reader = os.Stdin
	Output io.Writer = os.Stdout
)

// UserPrompt: Prompt for input and store it as variable.
func UserPrompt(prompt string, sensitive bool) (string, error) {
	return PromptWithOutput(prompt, sensitive, Output, "")
}

// UserPromptWithDefault: Prompt the user for input but if user don't put anything it will use default value
func UserPromptWithDefault(prompt, defaultValue string, sensitive bool) (string, error) {
	return PromptWithOutput(prompt, sensitive, Output, defaultValue)
}

// PromptPromptWithOutput: Prompt the user for input and gives the output
func UserPromptWithOutput(prompt string, sensitive bool, output io.Writer, defaultValue string) (string, error) {
	fmt.Fprintf(output, "%s: ", prompt)
	defer fmt.Fprintf(output, "\n")

	file, isFile := Input.(interface{ Fd() uintptr })

	if isFile && sensitive {
		if input, err := terminal.ReadPassword(int(file.Fd())); err != nil {
			return "", err
		} else {
			return strings.TrimSpace(string(input)), nil
		}
	} else {
		if value, err := bufio.NewReader(Input).ReadString('\n'); err != nil {
			return "", err
		} else {
			if value == "" || value == "\n" {
				return defaultValue, nil
			}
			return strings.TrimSpace(value), nil
		}
	}
}
