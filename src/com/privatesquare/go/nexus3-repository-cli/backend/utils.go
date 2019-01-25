package backend

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

/*
createBaseRequest create the base request for a HTTP request
@param method   string          http request method eg: GET, POST, etc
@param url      string          http request url
@param body     []byte          request body
@param user     m.AuthUser      User authentication details
@param verbose  boolean         prints verbose logs if set to true
@return *http.Request   HTTP base request
*/
func createBaseRequest(method, url string, body []byte) *http.Request {
	if SkipTLSVerification {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	req.SetBasicAuth(AuthUser.Username, AuthUser.Password)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	logError(err, "Error creating the request")

	if Verbose {
		fmt.Println("Request Url:", req.URL)
		fmt.Println("Request Headers:", req.Header)
		fmt.Println("Request Body:", req.Body)
	}

	return req
}

func createBaseRequest1(method, url string, body string) *http.Request {
	if SkipTLSVerification {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
	req, err := http.NewRequest(method, url, strings.NewReader(body))
	req.SetBasicAuth(AuthUser.Username, AuthUser.Password)
	req.Header.Set("Content-Type", "text/plain")
	logError(err, "Error creating the request")

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

func fileExists(fileName string) bool {
	if _, err := os.Stat(fileName); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func readFile(fileName string) string{
	var (
		data []byte
		err error
	)
	if fileExists(fileName) {
		data, err = ioutil.ReadFile(fileName)
		logError(err, "There was an error reading the file")
	}
	return string(data)
}
