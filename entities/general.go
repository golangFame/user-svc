package entities

type GenericResponse struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Meta       interface{} `json:"meta,omitempty"`
	StatusCode int         `json:"-"`
}
