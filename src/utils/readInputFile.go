package readInputFile

import (
	"fmt"
	"os"
	"strings"
)

func ReadInputFile(filePath string) []string {

	inputByte, inputErr := os.ReadFile(filePath)

	if inputErr != nil {
		fmt.Print(inputErr)
	}

	inputString := string(inputByte)
	lines := strings.Split(inputString, "\n")

	return lines
}
