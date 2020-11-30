// B"H

package structs

// IngressData contain necessary data to send to handler service.
type IngressData struct {
	Name   string `json:"name"`
	Host   string `json:"host"`
	Expose bool   `json:"expose"`
}
