package main

import (
	"fmt"
	"runtime"

	"github.com/xuri/excelize/v2"
)
type DocumentInfo struct{
	cellCoordinates []string
	variableNames []string
	variableContents []string	

}

// @param Filename has to include the name of the XLSX file ony XLSX is supported.
func createTemplateJsonFileFromXLSX(fileName string){
	// Get value from cell by given worksheet name and axis.
	file := loadExcelTemplateFile(fileName)
	regex := `(?:var_)`
	sheet := "Sheet1"
	cells, err := file.SearchSheet(sheet, regex, true)
	cellValues := make([]string, len(cells))
	var docInfo DocumentInfo
	checkForErrors(err)

	for i := 0; i < len(cells); i++ {
		cellValues[i], err = file.GetCellValue(sheet, cells[i])
		checkForErrors(err)
	}
	docInfo.cellCoordinates = cells
	docInfo.variableNames = cellValues
}
// Checks for errors, self explanatory. Shows stack trace if error is Encountered.
func checkForErrors(e error){
    if e == nil {
        return
    }
	pc := make([]uintptr, 10)  // at least 1 entry needed
    runtime.Callers(2, pc)
    f := runtime.FuncForPC(pc[0])
    file, line := f.FileLine(pc[0])
    fmt.Println(e)
    fmt.Printf("%s:%d %s\n", file, line, f.Name())

}

func loadExcelTemplateFile(filename string) *excelize.File{
    ExcelFile, error := excelize.OpenFile("../../templates/invoice"+filename+".xlsx")
    checkForErrors(error)
	    defer func() {
        // Close the spreadsheet.
        if err := ExcelFile.Close(); err != nil {
            fmt.Println(err)
        }
    }()
	return ExcelFile

}
func main() {
	// add some shit for it to use incoming data from the network.
	createTemplateJsonFileFromXLSX("invoice")
}