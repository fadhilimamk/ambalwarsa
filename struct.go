package main

// Response : ini comment
type Response struct {
	Data   interface{}   `json:"data,omitempty"`
	Errors []interface{} `json:"errors,omitempty"`
	Meta   interface{}   `json:"meta,omitempty"`
}
