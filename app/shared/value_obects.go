package shared

func AssertString(val any) string {
	if s, ok := val.(string); ok {
		return s
	}
	return ""
}
