// B"H

package utils

// credit https://blog.risingstack.com/golang-tutorial-for-nodejs-developers-getting-started/#nethttp

import (
	"io"
	"io/ioutil"
	"net/http"
	// logf "sigs.k8s.io/controller-runtime/pkg/log"
)

// MakePostRequest is
func MakePostRequest(url string, payloadBuf io.Reader) {

	method := "POST"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payloadBuf)

	if err != nil {
		log.Error(err, "error")
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Error(err, "error")
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error(err, "error")
	}
	log.Info(string(body))

}
