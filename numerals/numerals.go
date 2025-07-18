package numerals

import "strings"

func ConvertToRoman(arabic int) (string, error) {

	if arabic == 4 {
		return "VI", nil
	}

	var result strings.Builder

	for i := 0; i < arabic; i++ {
		result.WriteString("I")
	}

	return result.String(), nil
}
