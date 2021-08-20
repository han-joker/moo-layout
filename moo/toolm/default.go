package toolm

func StringDefault(value string, defValues ...string) string {
	if value != "" {
		return value
	}
	for _, v := range defValues {
		if v != "" {
			return v
		}
	}
	return ""
}

func Int64Default(value int64, defValues ...int64) int64 {
	if value != 0 {
		return value
	}
	for _, v := range defValues {
		if v != 0 {
			return v
		}
	}
	return 0
}