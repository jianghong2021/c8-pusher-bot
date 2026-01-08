package utils

import (
	"strings"
)

func EscapeMarkdownV2(text string) string {
	replacer := strings.NewReplacer(
		"_", "\\_",
		"*", "\\*",
		"[", "\\[",
		"]", "\\]",
		"(", "\\(",
		")", "\\)",
		"~", "\\~",
		"`", "\\`",
		">", "\\>",
		"#", "\\#",
		"+", "\\+",
		"-", "\\-",
		"=", "\\=",
		"|", "\\|",
		"{", "\\{",
		"}", "\\}",
		".", "\\.",
		"!", "\\!",
		"<", "\\<",
	)
	return replacer.Replace(text)
}

// 构建 Telegram MarkdownV2 格式的键值对列表
func BuildMarkdownV2List(headers []string, rows [][]string, leftPadding string) string {
	if len(headers) == 0 || len(rows) == 0 {
		return ""
	}

	var b strings.Builder

	for _, row := range rows {
		for i, header := range headers {
			if i < len(row) {
				escapedHeader := EscapeMarkdownV2(header)
				escapedValue := EscapeMarkdownV2(row[i])
				b.WriteString(leftPadding + escapedHeader + ": " + escapedValue + "\n")
			}
		}
		b.WriteString("\n")
	}

	return b.String()
}
