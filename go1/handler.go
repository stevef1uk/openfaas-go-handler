package function

import (
	handler "github.com/openfaas-incubator/go-function-sdk"
	//"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Handle a function invocation
func Handle(req handler.Request) (handler.Response, error) {
	var err error
	var body []byte
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

	if req.Host == "" {
		//req.Host = "http://gateway.openfaas:8080/function/env"
		//req.Host = "http://test4.openfaas:5000/v1/verysimple?id=1"
		req.Host = "http://test4.openfaas:5000/v1/verysimple?" + req.QueryString
	}
	log.Println("Host = " + req.Host)
	log.Println("Query String = " + req.QueryString)
	log.Println("Method = " + req.Method)
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
