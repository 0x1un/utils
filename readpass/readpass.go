package readpass

import (
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

func ReadPassword() (string, error) {
	fd := os.Stdin.Fd()
	pass, err := terminal.ReadPassword(int(fd))
	return string(pass), err
}
