package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"runtime"

	"github.com/xuri/excelize/v2"
)
type DocumentInfo struct{
	VariableNames []string `json:"variables"` 
	CellCoordinates []string `json:"coordinates"` 

}

// @param Filename has to include the name of the XLSX file ony XLSX is supported.
func createTemplateJsonFileFromXLSX(fileName string){
	// Get value from cell by given worksheet name and axis.
	file := loadExcelTemplateFile(fileName)
	regex := `(?:var_)`
	regExp := regexp.MustCompile(regex)
	sheet := "Sheet1"

	cellsCoordinates, err := file.SearchSheet(sheet, regex, true)

	varNames := make([]string, len(cellsCoordinates))

	var docInfo DocumentInfo
	checkForErrors(err)

	for i := 0; i < len(cellsCoordinates); i++ {
		varNames[i], err = file.GetCellValue(sheet, cellsCoordinates[i])
		checkForErrors(err)
		//remove the var_ from the field
		varNames[i] = cleanField(regExp.ReplaceAllString(varNames[i], ""))

	}
	
	docInfo.CellCoordinates = cellsCoordinates
	docInfo.VariableNames = varNames

	Fjson, err := json.MarshalIndent(docInfo,"","\t")
	checkForErrors(err)

	defer func(){
		err := os.WriteFile("../../templates/"+fileName+".json",Fjson,0644)
		checkForErrors(err)
		fmt.Println("template created: "+fileName+".json")
	}()

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
    ExcelFile, error := excelize.OpenFile("../../templates/"+filename+".xlsx")
    checkForErrors(error)
	    defer func() {
        // Close the spreadsheet.
        if err := ExcelFile.Close(); err != nil {
            fmt.Println(err)
        }
    }()
	return ExcelFile

}
func cleanField(field string) string{
	regExp := regexp.MustCompile(`(?: )`)
	return regExp.ReplaceAllString(field, "")
}

func main() {
	// add some shit for it to use incoming data from the network.
	createTemplateJsonFileFromXLSX("quote")
}