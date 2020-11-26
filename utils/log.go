// B"H

package utils

// credit https://blog.risingstack.com/golang-tutorial-for-nodejs-developers-getting-started/#nethttp

import (
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

var log = logf.Log.WithName("utils")

// LogDecleration is
func LogDecleration() {
	log.Info("In progress", "Now declaring log: ", "should never be used.")
}
