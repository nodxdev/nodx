package nodx

import "strings"

// escapeHTML escapes the input string to prevent XSS attacks.
func escapeHTML(input string) string {
	replacements := map[string]string{
		"&":  "&#38;",
		"<":  "&#60;",
		">":  "&#62;",
		"\"": "&#34;",
		"'":  "&#39;",
	}

	var builder strings.Builder

	for _, char := range input {
		strChar := string(char)
		if replacement, exists := replacements[strChar]; exists {
			builder.WriteString(replacement)
		} else {
			builder.WriteString(strChar)
		}
	}

	return builder.String()
}
