package excel_tools

import "fmt"

// columnNumberToName 将列号转换为 Excel 列名
func columnNumberToName(columnNumber int) string {
	columnName := ""
	for columnNumber > 0 {
		columnNumber-- // 调整为从0开始
		remainder := columnNumber % 26
		columnName = fmt.Sprintf("%c", 'A'+remainder) + columnName
		columnNumber = columnNumber / 26
	}
	return columnName
}

// GetExcelCellAddress 返回 Excel 单元格地址
// row: 行号(从1开始)
// column: 列号(从1开始)
func GetExcelCellAddress(row, column int) string {
	columnName := columnNumberToName(column)
	return fmt.Sprintf("%s%d", columnName, row)
}
