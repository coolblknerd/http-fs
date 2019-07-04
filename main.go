package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	// Since we've already deferenced the request we can now store this
	// request into memory to start messing around with the data. The rest
	// of the request ends up stored on disk in temp files. 10 << 20 ends
	// creating 10 MB of space. (10 << 20).
	//
	// How we get 10 MB from 10 << 20??
	//
	// "<<" is a bitwise operator that is going to shift our int of "10", 20
	// bits to the left which will give us a binary number of
	// "101000000000000000000000". Don't know if you know how to count
	// binary (It's cool if you don't, cause most don't. I used a calculator
	// honestly...lol) but the binary ends up translating to the decimal
	// 10485760 (binary to decimal, you follow?). This decimal ends up representing
	// how many bytes we're passing into the Parse function. Convert it to MB and
	// you get around 10 MB
	r.ParseMultipartForm(10 << 20)

	// FormFile returns the first file for the provided form key.
	file, handler, err := r.FormFile("test")
	if err != nil {
		fmt.Println("Error uploading the file")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	// This will read all of the data (as "bytes") into the memory.
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	// This takes the temp file just created and writes the data (in bytes)
	// to the temp file.
	tempFile.Write(fileBytes)

	// returns a message to let users know the file was uploaded
	fmt.Fprintf(w, "Successfully uploaded file")
}

func serverRequest() {
	http.HandleFunc("/upload", uploadFile)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func main() {
	fmt.Println("This is Reggie")
	serverRequest()
}
