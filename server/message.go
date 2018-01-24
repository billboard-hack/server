package server

// message represents a single message
type message struct {
	Name    string `json:"name,omitempty"`
	Message string `json:"message,omitempty"`
}
