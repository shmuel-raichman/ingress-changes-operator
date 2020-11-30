// B"H

package structs

// Config contain the environments variables that utils.config.go reads.
type Config struct {
	Port                    int
	Environment             string
	IngressesHandlerAddress string
	ExposeAnnotation        string
}
