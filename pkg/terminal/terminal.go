package terminal

import (
	"bufio"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func ReadString() string {
	readFromTerminalStr, _ := reader.ReadString('\n')
	readFromTerminalStr = strings.TrimSpace(readFromTerminalStr)

	return readFromTerminalStr
}
