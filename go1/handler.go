package function

import (
	"bytes"
	handler "github.com/openfaas-incubator/go-function-sdk"
	"io/ioutil"
	"log"
	"net/http"
)

// Handle a function invocation
func Handle(req handler.Request) (handler.Response, error) {
	var err error
	ret_msg := ""
	status := http.StatusOK

	log.Printf("In handler, req = %v\n", req)
	if checkSecretOk("secret-api-key", req) {
		//req.Host = "http://gateway.openfaas:8080/function/env"
		req.Host = "http://test4.openfaas:5000/v1/verysimple"
		switch req.Method {
		case "GET":
			ret_msg, err = handleGET(req)
		case "POST":
			ret_msg, err = handlePOST(req)
		default:
			ret_msg = "Error: unrecognised method: " + req.Method
			status = http.StatusMethodNotAllowed
		}
		if err != nil {
			status = http.StatusServiceUnavailable
		}
	} else {
		log.Println("API Request not validated")
		ret_msg = "API Key not present or valid "
		status = http.StatusForbidden
	}

	return handler.Response{
		Body:       []byte(ret_msg),
		StatusCode: status,
	}, err
}

func handleGET(req handler.Request) (string, error) {
	ret := ""
	var body []byte

	log.Println("In HandleGet, Host = " + req.Host)
	log.Println("Query String = " + req.QueryString)

	resp, err := http.Get(req.Host + "?" + req.QueryString)
	if err != nil {
		// handle error
		ret = "OOPs call failed " + err.Error()
	} else {
		body, err = ioutil.ReadAll(resp.Body)
		ret = string(body)
		if ret == "" {
			ret = "No Data Found"
		}
		defer resp.Body.Close()
	}
	return ret, err
}

func handlePOST(req handler.Request) (string, error) {
	ret := ""

	log.Println("In HandlePOST, Host = " + req.Host)
	log.Println("Body = " + string(req.Body))

	resp, err := http.Post(req.Host, "application/json", bytes.NewBuffer(req.Body))
	if err != nil {
		ret = "OOPs call failed " + err.Error() + " Response " + string(resp.StatusCode)
	}
	return ret, err
}

func getAPISecret(secretName string) (secretBytes []byte, err error) {
	// read from the openfaas secrets folder
	secretBytes, err = ioutil.ReadFile("/var/openfaas/secrets/" + secretName)
	if err != nil {
		// read from the original location for backwards compatibility with openfaas <= 0.8.2
		secretBytes, err = ioutil.ReadFile("/run/secrets/" + secretName)
		log.Println("Read Secret ok")
	}

	return secretBytes, err
}

func checkSecretOk(secretName string, req handler.Request) bool {
	ret := true

	log.Printf(" In checkSecretOk Header structure %v\n", req.Header)
	key := req.Header.Get("X-Api-Key")
	//log.Printf("API Key passed = %s\n", key)
	real_secret, err := getAPISecret("secret-api-key")
	if err == nil {
		log.Printf("comparing = %v to %v\n", []byte(key), real_secret)
		if !bytes.Equal([]byte(key), real_secret) {
			ret = false
		}
	}

	return ret
}
