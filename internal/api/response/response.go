package response

import "net/http"

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Code    string      `json:"code"`
	Data    interface{} `json:"data,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
}



// Success response
func Success(
	message string,
	data interface{},
	meta interface{},
) (int, APIResponse) {

	return http.StatusOK, APIResponse{
		Success: true,
		Message: message,
		Code:    "SUCCESS",
		Data:    data,
		Meta:    meta,
	}
}

// Created response
func Created(
	message string,
	data interface{},
) (int, APIResponse) {

	return http.StatusCreated, APIResponse{
		Success: true,
		Message: message,
		Code:    "CREATED",
		Data:    data,
	}
}

// Error response
func Error(
	status int,
	message string,
	code string,
) (int, APIResponse) {

	return status, APIResponse{
		Success: false,
		Message: message,
		Code:    code,
	}
}
