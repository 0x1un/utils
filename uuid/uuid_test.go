package uuid

import (
	"fmt"
	"testing"
)

func TestMakeUUID(t *testing.T) {
	u := NewUUID()
	fmt.Println(u.String())
}
