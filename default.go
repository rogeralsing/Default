package godefault

func String(value string, defaultValue string) string {
	if value == "" {
		return defaultValue
	}
	return value
}

func Int(value int, defaultValue int) int {
	if value == 0 {
		return defaultValue
	}
	return value
}
