package server

// message represents a single message
type message struct {
<<<<<<< HEAD
	Name      string
	Message   string
	When      time.Time
=======
	Name    string `json:"name,omitempty"`
	Message string `json:"message,omitempty"`
>>>>>>> fb0285c... commit
}
