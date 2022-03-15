package determine

import (
	"fmt"
	"testing"
)

func TestDirExists(t *testing.T) {
	fmt.Println(DirExists("../diff"))
}

func TestFileExists(t *testing.T) {
	fmt.Println(FileExists("../diff/diff.go"))
}
