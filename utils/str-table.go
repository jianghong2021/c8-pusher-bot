package utils

import (
	"strings"
	"unicode/utf8"
)

// 格式化表格为 Telegram MarkdownV2
// headers: 表头
// rows: 数据行
// align: 每列的对齐方式 ("left", "center", "right")
func FormatMarkdownTable(headers []string, rows [][]string, align []string) string {
	if len(headers) == 0 || len(rows) == 0 {
		return ""
	}

	// 初始化对齐方式，默认为左对齐
	if len(align) < len(headers) {
		align = make([]string, len(headers))
		for i := range align {
			align[i] = "left"
		}
	}

	// 计算每列的最大宽度
	colWidths := make([]int, len(headers))

	// 统计表头宽度
	for i, header := range headers {
		escapedHeader := EscapeMarkdownV2(header)
		colWidths[i] = utf8.RuneCountInString(escapedHeader)
	}

	// 统计数据行宽度
	for _, row := range rows {
		for i := 0; i < len(headers) && i < len(row); i++ {
			escapedCell := EscapeMarkdownV2(row[i])
			cellWidth := utf8.RuneCountInString(escapedCell)
			if cellWidth > colWidths[i] {
				colWidths[i] = cellWidth
			}
		}
	}

	var builder strings.Builder

	// 添加表格边框
	builder.WriteString("```\n")

	// 构建表头
	for i, header := range headers {
		escapedHeader := EscapeMarkdownV2(header)
		builder.WriteString("| ")
		switch align[i] {
		case "center":
			padding := colWidths[i] - utf8.RuneCountInString(escapedHeader)
			leftPad := padding / 2
			rightPad := padding - leftPad
			builder.WriteString(strings.Repeat(" ", leftPad))
			builder.WriteString(escapedHeader)
			builder.WriteString(strings.Repeat(" ", rightPad))
		case "right":
			spaces := colWidths[i] - utf8.RuneCountInString(escapedHeader)
			builder.WriteString(strings.Repeat(" ", spaces))
			builder.WriteString(escapedHeader)
		default: // left
			builder.WriteString(escapedHeader)
			spaces := colWidths[i] - utf8.RuneCountInString(escapedHeader)
			builder.WriteString(strings.Repeat(" ", spaces))
		}
		builder.WriteString(" ")
	}
	builder.WriteString("|\n")

	// 构建分隔线
	for i, width := range colWidths {
		builder.WriteString("|-")
		switch align[i] {
		case "left":
			builder.WriteString(strings.Repeat("-", width))
		case "center":
			builder.WriteString(strings.Repeat("-", width))
		case "right":
			builder.WriteString(strings.Repeat("-", width))
		}
		builder.WriteString("-")
	}
	builder.WriteString("|\n")

	// 构建数据行
	for _, row := range rows {
		for i := 0; i < len(headers); i++ {
			var cellValue string
			if i < len(row) {
				cellValue = row[i]
			}

			escapedCell := EscapeMarkdownV2(cellValue)
			builder.WriteString("| ")

			switch align[i] {
			case "center":
				padding := colWidths[i] - utf8.RuneCountInString(escapedCell)
				leftPad := padding / 2
				rightPad := padding - leftPad
				builder.WriteString(strings.Repeat(" ", leftPad))
				builder.WriteString(escapedCell)
				builder.WriteString(strings.Repeat(" ", rightPad))
			case "right":
				spaces := colWidths[i] - utf8.RuneCountInString(escapedCell)
				builder.WriteString(strings.Repeat(" ", spaces))
				builder.WriteString(escapedCell)
			default: // left
				builder.WriteString(escapedCell)
				spaces := colWidths[i] - utf8.RuneCountInString(escapedCell)
				builder.WriteString(strings.Repeat(" ", spaces))
			}
			builder.WriteString(" ")
		}
		builder.WriteString("|\n")
	}

	builder.WriteString("```")
	return builder.String()
}

// 高级表格格式化，支持可选的样式
func FormatAdvancedTable(headers []string, rows [][]string, options ...TableOption) string {
	opts := defaultTableOptions()
	for _, opt := range options {
		opt(&opts)
	}

	return FormatMarkdownTableWithOptions(headers, rows, opts)
}

// 表格选项
type TableOptions struct {
	Align      []string
	Border     bool
	MaxRows    int
	Truncate   bool
	Footer     []string
	HideHeader bool
}

func defaultTableOptions() TableOptions {
	return TableOptions{
		Align:      []string{},
		Border:     true,
		MaxRows:    0, // 0表示不限制
		Truncate:   true,
		HideHeader: false,
	}
}

type TableOption func(*TableOptions)

func WithAlign(align []string) TableOption {
	return func(o *TableOptions) {
		o.Align = align
	}
}

func WithBorder(border bool) TableOption {
	return func(o *TableOptions) {
		o.Border = border
	}
}

func WithMaxRows(maxRows int) TableOption {
	return func(o *TableOptions) {
		o.MaxRows = maxRows
	}
}

func WithFooter(footer []string) TableOption {
	return func(o *TableOptions) {
		o.Footer = footer
	}
}

func WithHideHeader(hide bool) TableOption {
	return func(o *TableOptions) {
		o.HideHeader = hide
	}
}

// 带有选项的表格格式化
func FormatMarkdownTableWithOptions(headers []string, rows [][]string, opts TableOptions) string {
	if len(headers) == 0 || len(rows) == 0 {
		return "No data available"
	}

	// 限制行数
	if opts.MaxRows > 0 && len(rows) > opts.MaxRows && opts.Truncate {
		rows = rows[:opts.MaxRows]
	}

	// 添加尾部行
	if len(opts.Footer) > 0 {
		newRows := make([][]string, len(rows))
		copy(newRows, rows)
		newRows = append(newRows, opts.Footer)
		rows = newRows
	}

	// 如果不显示表头，但需要计算宽度
	displayHeaders := headers
	if opts.HideHeader {
		displayHeaders = make([]string, len(headers))
		for i := range displayHeaders {
			displayHeaders[i] = "" // 用空字符串代替
		}
	}

	// 生成表格
	table := FormatMarkdownTable(displayHeaders, rows, opts.Align)

	// 添加行数信息
	if opts.MaxRows > 0 && len(rows) >= opts.MaxRows && opts.Truncate {
		table += "\n\n_*Showing " + string(rune(opts.MaxRows)) + " of " + string(rune(len(rows))) + " rows_*"
	}

	return table
}
