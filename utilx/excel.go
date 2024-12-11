package utilx

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

func Xlsx(data [][]interface{}, sheetName string) (*excelize.File, error) {
	f := excelize.NewFile()

	index := 0
	oldName := f.GetSheetName(index)
	err := f.SetSheetName(oldName, sheetName)
	if err != nil {
		return nil, fmt.Errorf("SetSheetName failed: %s", err.Error())
	}

	// 遍历二维字符串数组并写入数据到Excel文件
	for rowIdx, rowData := range data {
		for colIdx, cellData := range rowData {
			cell, err := excelize.CoordinatesToCellName(colIdx+1, rowIdx+1) // Excel索引从1开始
			if err != nil {
				return nil, fmt.Errorf("CoordinatesToCellName failed: %s", err.Error())
			}
			err = f.SetCellValue(sheetName, cell, cellData)
			if err != nil {
				return nil, fmt.Errorf("SetCellValue failed: %s", err.Error())
			}
		}
	}

	return f, nil
}

// XlsxV2 流式写入，节约内存
func XlsxV2(data [][]interface{}, sheetName string) (*excelize.File, error) {
	f := excelize.NewFile()
	index := 0
	oldName := f.GetSheetName(index)
	err := f.SetSheetName(oldName, sheetName)

	if err != nil {
		return nil, fmt.Errorf("SetSheetName failed: %s", err.Error())
	}

	writer, err := f.NewStreamWriter(sheetName)
	if err != nil {
		return nil, fmt.Errorf("f.NewStreamWriter failed, err: %v", err)
	}

	// 遍历二维字符串数组并写入数据到Excel文件
	for rowIdx, rowData := range data {
		cell, err := excelize.CoordinatesToCellName(1, rowIdx+1) // Excel索引从1开始
		if err != nil {
			return nil, fmt.Errorf("CoordinatesToCellName failed: %s", err.Error())
		}
		err = writer.SetRow(cell, rowData)
		if err != nil {
			return nil, fmt.Errorf("writer.SetRow failed, err: %v", err)
		}
	}

	err = writer.Flush()

	if err != nil {
		return nil, fmt.Errorf("writer.Flush failed, err: %v", err)
	}

	return f, nil
}
