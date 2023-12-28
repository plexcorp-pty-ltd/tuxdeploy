package termio

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

func PromptPassword(question string) string {
	fmt.Fprint(os.Stderr, GetColorText(TEXT_SPECIAL_LIGHT, TEXT_SPECIAL_DARK).Render(question+" "))
	fmt.Println()

	password, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("Error reading password:", err)
		os.Exit(1)
	}

	return strings.TrimSpace(string(password))
}
