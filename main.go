package main

import (

	"log"
    function "github.com/stevef1uk/test2/go1"
    handler "github.com/openfaas-incubator/go-function-sdk"
)


func main() {
	log.Print("Hello world sample started.")

	var  req = new(handler.Request)

    req.Host = "http://gateway.openfaas:8080/function/env"




	resp, err := function.Handle(*req)
	if err != nil {
	    log.Print("That didn't work!")
	}
	log.Print("Body = " + string(resp.Body))
}
