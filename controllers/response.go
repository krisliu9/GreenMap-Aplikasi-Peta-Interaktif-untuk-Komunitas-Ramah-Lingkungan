package controllers

type Response struct {
	Status     bool        `json:"status"`
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Token      interface{} `json:"token,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}
