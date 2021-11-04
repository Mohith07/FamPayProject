package model

// Response struct
type Response struct {
	Message      string      `json:"message,omitempty"`
	ErrorMessage interface{} `json:"error_message,omitempty"`
	Body         interface{} `json:"body,omitempty"`
	Page		int			  `json:"page"`
}

