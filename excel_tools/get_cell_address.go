package excel_tools

import (
	"fmt"
	"github.com/spf13/cast"
	"log"
	"strings"
	"unicode"
)

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

// columnNameToNumber 将 Excel 列名转换为列号
func columnNameToNumber(columnName string) int {
	columnNumber := 0
	for i := 0; i < len(columnName); i++ {
		columnNumber = columnNumber*26 + int(columnName[i]-'A'+1)
	}
	return columnNumber
}

// GetExcelCellAddress 返回 Excel 单元格地址
// row: 行号(从1开始)
// column: 列号(从1开始)
func GetExcelCellAddress(row, column int) string {
	columnName := columnNumberToName(column)
	return fmt.Sprintf("%s%d", columnName, row)
}

// GetExcelCellRowColumn 返回 Excel 单元格的行号和列号
// cellAddress: 单元格地址
// return: 行号, 列号
func GetExcelCellRowColumn(cellAddress string) (int, int) {
	var row, column strings.Builder
	for _, r := range cellAddress {
		if unicode.IsLetter(r) {
			column.WriteRune(r)
		} else if unicode.IsDigit(r) {
			row.WriteRune(r)
		} else {
			log.Println("invalid cell address")
			return -1, -1
		}
	}

	if column.Len() == 0 || row.Len() == 0 {
		log.Println("invalid cell address")
		return -1, -1
	}
	columnInt := columnNameToNumber(column.String())
	rowInt := cast.ToInt(row.String())
	return rowInt, columnInt
}
