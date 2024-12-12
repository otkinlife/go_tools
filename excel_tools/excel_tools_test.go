package excel_tools

import "testing"

func TestGetExcelCellAddress(t *testing.T) {
	addr := GetExcelCellAddress(1, 2)
	t.Logf("GetExcelCellAddress() succeeded, addr: %s", addr)
	return
}

func TestGetExcelCellRowColumn(t *testing.T) {
	row, column := GetExcelCellRowColumn("B1")
	t.Logf("GetExcelCellRowColumn() succeeded, row: %d, column: %d", row, column)
	return
}
