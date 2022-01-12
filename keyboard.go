// Package keyboard reads user input from the keyboard
package keyboard

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// GetFloat needs a floating-point number from the kyboard.
// It returns the number read and any error encountered
func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)

	if err != nil {
		return 0, err
	}

	return number, nil
}
