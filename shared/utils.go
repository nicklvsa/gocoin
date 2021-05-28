package shared

func GetPointerToString(s string) *string {
	return &s
}

func IsStringPtrNilOrEmpty(inputs ...*string) bool {
	for _, input := range inputs {
		if input == nil || *input == "" {
			return true
		}
	}

	return false
}
