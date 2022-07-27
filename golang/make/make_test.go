package make

import (
	"fmt"
	"testing"
)

func TestMake(t *testing.T) {
	s := []string{"go", "python", "shell"}
	fmt.Println(s)
	s = make([]string, 1)

	s[0] = "java"
	fmt.Println(s)
}
