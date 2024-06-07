package views

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetTextInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}