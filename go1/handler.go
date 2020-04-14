package function

import (
	//"fmt"
	"io/ioutil"
	"log"
	"net/http"

	handler "github.com/openfaas-incubator/go-function-sdk"
)

// Handle a function invocation
func Handle(req handler.Request) (handler.Response, error) {
	var err error
	var body []byte
	ret_msg := "Hello world, from Steve & Sarah"
	//message := fmt.Sprintf("Hello world, from Steve & Sarah  the input was: %s", string(req.Body))
	log.Print("In handler")
	if req.Host == "" {
		req.Host = "https://gateway.openfaas:8080/function/env"
	}
	log.Print("Attempting to call URL " + req.Host)
	resp, err := http.Get(req.Host)
	if err != nil {
		// handle error
		ret_msg = "OOPs call failed " + err.Error()
		body = []byte(ret_msg)
	} else {
		body, err = ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
	}

	return handler.Response{
		Body:       []byte(body),
		StatusCode: http.StatusOK,
	}, err
}
