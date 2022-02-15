package service

import (
	"goginsoap/soapHandler"
)

type Service struct{}

type Stock struct {
	ID         string
	Name       string
	Code       string
	Address    string
	PostalCode string
}

// RetrieveStock - retrieves comments by filters
func RetrieveStock(request soapHandler.Request) (*soapHandler.Response, error) {

	soapRequest := request
	response, err := soapHandler.CallSOAPClientSteps(&soapRequest)

	if err != nil {
		return response, err
	}

	return response, nil
}
