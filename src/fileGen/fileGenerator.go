package main

import (
	"fmt"
	"runtime"
	"strconv"

	"github.com/xuri/excelize/v2"
)


type DocumentInfo struct{
	cellCoordinates []string
	variableNames []string
	variableContents []string	

}
type DocumentContentInfo struct {
	variableNames []string
	variableContents []string
}
func main() {
	docInfo := getEditableData(loadTemplate())
	
	createNewFile(addVariableContent(docInfo))


}

func addVariableContent(docInfo DocumentInfo) DocumentInfo {
	var info DocumentContentInfo

	names := []string{"var_description","var_total"}
	values := []string{"Caya butishi 205","40000"}
	info.variableNames,docInfo.variableNames = names,names
	info.variableContents,docInfo.variableContents = values,values

	return docInfo
}


func loadTemplate() *excelize.File{
    template, error := excelize.OpenFile("../templates/invoice.xlsx")
    checkForErrors(error, "loadTemplate")
	    defer func() {
        // Close the spreadsheet.
        if err := template.Close(); err != nil {
            fmt.Println(err)
        }
    }()
	return template

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

func createNewFile(docInfo DocumentInfo){
	f := loadTemplate()

	// Create a new sheet.
	// Set value of a cell.
	
	fmt.Println(docInfo)
	for i:=0; i<len(docInfo.cellCoordinates); i++{
		val := docInfo.variableContents[i]
		integer, err := strconv.ParseInt(val,10,64)
		if err == nil {
			f.SetCellValue("Sheet1", docInfo.cellCoordinates[i], integer)
			
		}else{
			f.SetCellValue("Sheet1", docInfo.cellCoordinates[i], val)
		}
	}

	// Set active sheet of the workbook
	f.SetActiveSheet(0)
	// Save spreadsheet by he given path.
	if err := f.SaveAs("Book1.xlsx"); err == nil {
		fmt.Println(err)
	}

}

