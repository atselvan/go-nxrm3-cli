package backend

import (
	"bytes"
	"com/privatesquare/go/nexus3-repository-cli/model"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
)

/*
logError prints error
@param err error  error details
@return void
*/
func logError(err error, errorMessage string) {
	if err != nil {
		log.Println(errorMessage)
		log.Fatal(err)
	}
}

func logJsonMarshalError(err error, funcName string) {
	if err != nil {
		log.Println(fmt.Sprintf("%s : %s", funcName, jsonMarshalError))
		os.Exit(1)
	}
}

func logJsonUnmarshalError(err error, funcName string) {
	if err != nil {
		log.Println(fmt.Sprintf("%s : %s", funcName, jsonUnmarshalError))
		os.Exit(1)
	}
}

/*
createBaseRequest create the base request for a HTTP request
@param method   string          http request method eg: GET, POST, etc
@param url      string          http request url
@param body     []byte          request body
@param user     m.AuthUser      User authentication details
@param verbose  boolean         prints verbose logs if set to true
@return *http.Request   HTTP base request
*/
func createBaseRequest(method, url string, requestBody model.RequestBody) *http.Request {
	if SkipTLSVerification {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
	var (
		req *http.Request
		err error
	)
	if requestBody.Json != nil {
		req, err = http.NewRequest(method, url, bytes.NewBuffer(requestBody.Json))
		logError(err, "Error creating the request")
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json")
	} else if requestBody.Text != "" {
		req, err = http.NewRequest(method, url, strings.NewReader(requestBody.Text))
		req.Header.Set("Content-Type", "text/plain")
		logError(err, "Error creating the request")
	} else {
		req, err = http.NewRequest(method, url, nil)
		logError(err, "Error creating the request")
	}
	req.SetBasicAuth(AuthUser.Username, AuthUser.Password)
	if Verbose {
		fmt.Println("Request Url:", req.URL)
		fmt.Println("Request Headers:", req.Header)
		fmt.Println("Request Body:", req.Body)
	}
	return req
}

/*
httpRequest makes a request to the remote server via a proxy server
@param user     m.AuthUser      User authentication details
@param req      *http.Request   HTTP base request
@param verbose  boolean         prints verbose logs if set to true
@return []byte  response body
@return string  response status
*/
func httpRequest(req *http.Request) ([]byte, string) {
	client := &http.Client{}
	resp, err := client.Do(req)
	logError(err, "There was a problem in making the request")

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	logError(err, "There was a problem reading the response body")

	if Verbose {
		fmt.Println("Response Headers:", resp.Header)
		fmt.Println("Response Status:", resp.Status)
		fmt.Println("Response Body:", string(respBody))
	}
	return respBody, resp.Status
}

// fileExists - Checks if a file exists
// @fileName: name or path to the file
func fileExists(fileName string) bool {
	if _, err := os.Stat(fileName); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// createFile - creates a new file
// @fileName: name or path to the file
func createFile(fileName string) *os.File {
	f, err := os.Create(fileName)
	logError(err, fmt.Sprintf("There was an error while creating the file %s", fileName))
	return f
}

// readFile - reads a file and returns the data in string format.
// The function checks if the file exists or not before reading the file
// @fileName: name or path to the file
func readFile(fileName string) string {
	var (
		data []byte
		err  error
	)
	if fileExists(fileName) {
		data, err = ioutil.ReadFile(fileName)
		logError(err, "There was an error reading the file.")
	} else {
		log.Printf("File %q was not found", fileName)
		os.Exit(1)
	}
	if string(data) == "" {
		log.Printf("The file %q is empty", fileName)
		os.Exit(1)
	}
	return string(data)
}

// readStringSliceFromFile - reads the data from a file and inputs every new line into a slice and returns the slice
// @fileName: name or path to the file
func readStringSliceFromFile(fileName string) []string {
	data := readFile(fileName)
	return strings.Split(strings.TrimSpace(data), "\n")
}

// writeFile - writes data to a file.
// If the file does not exists, a new file will be created
// if the file exists then the file will be overwritten with the new data
// @fileName: name or path to the file
// @data: the data in []byte format
func writeFile(fileName string, data []byte) {
	err := ioutil.WriteFile(fileName, data, 0644)
	logError(err, "There was an error writing to the file.")
}

// writeFile - writes a string slice to a file. Every new line will be written in a new line of the file
// If the file does not exists, a new file will be created
// if the file exists then the file will be overwritten with the new data
// @fileName: name or path to the file
// @slice: a string slice
func writeStringSliceToFile(fileName string, slice []string) {
	f := createFile(fileName)
	sep := "\n"
	for _, line := range slice {
		if _, err := f.WriteString(line + sep); err != nil {
			panic(err)
		}
	}
}

func printStringSlice(slice []string) {
	for _, s := range slice {
		fmt.Println(s)
	}
}

func entryExists(slice []string, entry string) bool {
	for i := 0; i < len(slice); i++ {
		if slice[i] == entry {
			return true
		}
	}
	return false
}

func getfuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

func toLower(s string) string {
	return strings.ToLower(s)
}

func getSliceIndex(slice []string, entry string) int {
	for i := 0; i < len(slice); i++ {
		if slice[i] == entry {
			return i
		}
	}
	return -1
}

func removeEntryFromSlice(slice []string, entry string) []string {
	i := getSliceIndex(slice, entry)
	if i == -1 {
		log.Printf("The entry %s does not exist, hence it cannot be removed", entry)
		os.Exit(1)
	}
	return append(slice[:i], slice[i+1:]...)
}
