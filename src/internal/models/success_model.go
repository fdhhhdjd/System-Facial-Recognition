package models

type SuccessResponse struct {
	Message          string      `json:"message"`
	Status           int         `json:"status"`
	ReasonStatusCode string      `json:"reason_status_code"`
	Option           interface{} `json:"option,omitempty"`
	Metadata         interface{} `json:"metadata,omitempty"`
}
