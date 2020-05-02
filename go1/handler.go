package function

import (
	"bytes"
	handler "github.com/openfaas-incubator/go-function-sdk"
	//"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Handle a function invocation
func Handle(req handler.Request) (handler.Response, error) {
	var err error
	ret_msg := "Hello world, from Steve & Sarah"
	//message := fmt.Sprintf("Hello world, from Steve & Sarah  the input was: %s", string(req.Body))
	log.Printf("In handler, req = %v\n", req)
	/*)
	if  len(req.QueryString) == 0 {
		log.Println("Empty Query String")
		return handler.Response{
			Body:       []byte(body),
			StatusCode: http.StatusBadRequest,
		}, err
	}

	id := req.QueryString[
	if req.Host == "" {
		req.Host = "http://gateway.openfaas:8080/function/env"
		//req.Host = "http://test4.default?id=1"
	}
	params := url.Values{}
	params.Add("id", string(id))
	req.Host = req.Host + params.Encode()
	*/
	req.Host = "http://test4.openfaas:5000/v1/verysimple"
	switch req.Method {
	case "GET":
		ret_msg, err = handleGET(req)
	case "POST":
		ret_msg, err = handlePOST(req)
	default:
		ret_msg = "Error: unrecognised method: " + req.Method
	}

	return handler.Response{
		Body:       []byte(ret_msg),
		StatusCode: http.StatusOK,
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
		defer resp.Body.Close()
	}
	return ret, err
}

func handlePOST(req handler.Request) (string, error) {
	ret := ""
	var body []byte

	log.Println("In HandlePOST, Host = " + req.Host)
	log.Println("Query String = " + req.QueryString)
	body = []byte(req.QueryString)

	resp, err := http.Post(req.Host, "application/json", bytes.NewBuffer(body))
	if err != nil {
		// handle error
		ret = "OOPs call failed " + err.Error()
	} else {
		body, err = ioutil.ReadAll(resp.Body)
		ret = string(body)
		defer resp.Body.Close()
	}
	return ret, err
}
