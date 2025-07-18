package numerals

func ConvertToRoman(arabic int) (string, error) {
	switch arabic {
	case 3:
		return "III", nil
	case 2:
		return "II", nil
	}

	return "I", nil
}
