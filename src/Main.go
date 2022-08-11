package main

import (
	"fmt"
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

func checkForErrors(e error, where string){
    if e == nil {
        return
    }
	fmt.Println("THIS IS AN ERROR YOU FUCKED UP around here"+ where )
    fmt.Println(e)
}
func getEditableData(file *excelize.File) DocumentInfo{
	// Get value from cell by given worksheet name and axis.
	regex := `(?:var_)` 
	sheet := "Sheet1"
    cells, err := file.SearchSheet(sheet, regex, true)
	cellValues := make([]string,len(cells))
	var docInfo DocumentInfo
	checkForErrors(err, "getEditableData")

	for i:= 0; i< len(cells); i++ {
		cellValues[i], err = file.GetCellValue(sheet,cells[i])
		checkForErrors(err, "getEditableData")
	}
	docInfo.cellCoordinates = cells
	docInfo.variableNames = cellValues
    return docInfo
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

