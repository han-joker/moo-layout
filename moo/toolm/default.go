package toolm

func StringDefault(value, defValue string) string {
	if value == "" {
		return defValue
	}
	return value
}
