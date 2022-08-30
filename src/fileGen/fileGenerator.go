package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"unicode"

	"log"
	"net/http"
	"runtime"
	"strconv"

	"github.com/xuri/excelize/v2"
)

// this is the expected structure of the incoming data that will come from firebase
type JsonPayload struct {
	Docnumber   string   `json: "docNumber"`
	Docname     string   `json: "docName"`
	Doctype     string   `json: "docType"`
	Variables   []string `json:"variables"`
	Coordinates []string `json:"coordinates"`
	Data        []string `json:"data"`
}

func loadTemplate() *excelize.File {
	template, error := excelize.OpenFile("template.xlsx")
	checkForErrors(error)
	defer func() {
		// Close the spreadsheet.
		if err := template.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	return template

}

// Checks for errors, self explanatory. Shows stack trace if error is Encountered.
func checkForErrors(e error) {
	if e == nil {
		return
	}
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	fmt.Println(e)
	fmt.Printf("%s:%d %s\n", file, line, f.Name())

}

func createNewFile(docInfo JsonPayload) (string, bool) {
	f := loadTemplate()

	// Create a new sheet.
	// Set value of a cell.

	fmt.Println(docInfo)
	for i := 0; i < len(docInfo.Coordinates); i++ {
		val := docInfo.Data[i]
		integer, err := strconv.ParseInt(val, 10, 64)
		if err == nil {
			f.SetCellValue("Sheet1", docInfo.Coordinates[i], integer)

		} else {
			f.SetCellValue("Sheet1", docInfo.Coordinates[i], val)
		}
	}

	// Set active sheet of the workbook
	f.SetActiveSheet(0)

	r := []rune(docInfo.Doctype)
	r[0] = unicode.ToUpper(r[0])
	doctypeCapitalCase := string(r)

	filename := doctypeCapitalCase + " " + docInfo.Docnumber + " " + docInfo.Docname + ".xlsx"
	// Save spreadsheet by he given path.
	if err := f.SaveAs(filename); err != nil {
		f.Close()
		fmt.Println(err)
		return filename, false

	}
	f.Close()
	return filename, true

}
func main() {

	// handle the file Upload and the data
	http.HandleFunc("/createFile", func(res http.ResponseWriter, req *http.Request) {

		//get the json payload
		var jsonPayload JsonPayload

		err := json.NewDecoder(req.Body).Decode(&jsonPayload)

		if err != nil {
			panic(err)
		}

		//get the file from the db
		fileUrl := ""

		if jsonPayload.Doctype == "invoice" {
			fileUrl = "https://firebasestorage.googleapis.com/v0/b/excelzior-83abd.appspot.com/o/SMConstruction%2Finvoice.xlsx?alt=media&token=5219766b-9789-43dd-8fe5-d5b6350d14b1"
		}

		fileErr := DownloadFile("template.xlsx", fileUrl)

		if fileErr != nil {
			panic(fileErr)
		}

		newFile, success := createNewFile(jsonPayload)

		if !success {

			fmt.Println("an error occurred")
		}

		//go to next step send file via http
		defer func() {
			data, err := os.ReadFile(newFile)
			if err != nil {
				panic(err)
			}
			res.Write([]byte(data))

		}()

	})

	serveServer()

}

func serveServer() {
	port := "6969"
	//serve the server
	fmt.Printf("Starting server at port " + port + "\n")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
