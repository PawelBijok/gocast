package utils

import "fmt"

func LeftPad(value string, pad int) string {
	pad = pad - len(value)
	padStr := ""
	for i := 0; i < pad; i++ {
		padStr += " "

	}
	return fmt.Sprintf("%s%s", padStr, value)
}
