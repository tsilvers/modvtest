package show

import (
	"fmt"
	"strings"
)

func Indent(spaces int, msg string) {
	indent := strings.Repeat(" ", spaces)
	fmt.Printf("%s%s (1)\n", indent, msg)
}
