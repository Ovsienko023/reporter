package safe

func String(str *string) string {
	if str == nil {
		return ""
	}
	return *str
}
