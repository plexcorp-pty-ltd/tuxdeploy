package termio

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

func PromptPassword(question string) string {
	fmt.Fprint(os.Stderr, GetColorText("#FF3333").Render(question+" "))
	fmt.Println()

	password, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("Error reading password:", err)
		os.Exit(1)
	}

	return strings.TrimSpace(string(password))
}
