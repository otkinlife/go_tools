package excel_tools

import "testing"

func TestGetExcelCellAddress(t *testing.T) {
	addr := GetExcelCellAddress(1, 2)
	t.Logf("GetExcelCellAddress() succeeded, addr: %s", addr)
	return
}
