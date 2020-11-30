// B"H

package utils

// credit https://blog.risingstack.com/golang-tutorial-for-nodejs-developers-getting-started/#nethttp

import (
	"fmt"
	"github.com/smuel1414/ingresses-changes/structs"
	"os"
	"strconv"
)

// ReadConfig getting series of environments variables to golang struct and checking if env variables is empty giving value.
func ReadConfig() structs.Config {

	portString := os.Getenv("PORT")
	if portString == "" {
		portString = "8000"
	}

	environment := os.Getenv("ENVIRONMENT")
	if environment == "" {
		environment = "local"
	}

	ingressesHandlerAddress := os.Getenv("INGRESSES_HANDLER_ADDRESS")
	if ingressesHandlerAddress == "" {
		ingressesHandlerAddress = "http://basic-http-server:8000"
	}

	exposeAnnotation := os.Getenv("EXPOSE_ANNOTATION")
	if exposeAnnotation == "" {
		exposeAnnotation = "expose.dns"
	}

	// Convert PORT envinroment variable to int.
	port, err := strconv.Atoi(portString)
	if err != nil {
		panic(fmt.Sprintf("Could not parse %s to int", portString))
	}

	return structs.Config{
		Port:                    port,
		Environment:             environment,
		IngressesHandlerAddress: ingressesHandlerAddress,
		ExposeAnnotation:        exposeAnnotation,
	}
}
