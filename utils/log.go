// B"H

package utils

// credit https://blog.risingstack.com/golang-tutorial-for-nodejs-developers-getting-started/#nethttp

import (
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

// Set a global logger for the memcached request. Each log record produced
// by this logger will have an identifier containing "request".
// These names are hierarchical; the name attached to request log statements
// will be "operator-sdk.request" because SDKLog has name
// "operator-sdk".
var log = logf.Log.WithName("utils")

// LogDecleration is
func LogDecleration() {
	log.Info("In progress", "Now declaring log: ", "should never be used.")
}
