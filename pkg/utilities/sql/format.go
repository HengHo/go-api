package sql

func FormatStringValue(data string) string {
	return "'" + data + "'"
}

func FormatRecord(data string) string {
	return "(" + data + ")"
}
