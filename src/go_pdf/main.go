package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

func main() {

	// handle the file Upload and the data
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {

		data, err := io.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}

		createFile(data)

		cmd := exec.Command("libreoffice", "--headless", "--convert-to pdf:writer_pdf_Export", "./file.xlsx")

		cmd.Run()

		defer func() {
			isDone := false
			for !isDone {
				if _, err := os.Stat("/file.pdf"); errors.Is(err, os.ErrNotExist) {
					time.Sleep(2 * time.Second)
				} else {
					isDone = true
				}
			}
			fileBytes, err := os.ReadFile("file.pdf")
			if err != nil {
				panic(err)
			}
			res.WriteHeader(http.StatusOK)
			res.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
			res.Write(fileBytes)
		}()

	})

	serveServer()

}

func serveServer() {
	port := "6970"
	//serve the server
	fmt.Printf("Starting server at port " + port + "\n")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
func createFile(fileBytes []byte) {
	f, err := os.Create("file.xlsx")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	amount, err2 := f.Write(fileBytes)
	fmt.Println(amount)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("done")
}
