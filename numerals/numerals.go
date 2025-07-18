package numerals

import "strings"

func ConvertToRoman(arabic int) (string, error) {

	var result strings.Builder

	for i := 0; i < arabic; i++ {
		result.WriteString("I")
	}

	return result.String(), nil
}
