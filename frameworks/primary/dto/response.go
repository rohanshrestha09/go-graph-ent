package dto

import (
	"math"
	"time"
)

type Response[T any] struct {
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func NewResponse[T any](message string, data T) *Response[T] {
	return &Response[T]{
		Message: message,
		Data:    data,
	}
}

type PaginatedResponse[T any] struct {
	Message     string `json:"message"`
	Result      []T    `json:"result"`
	Count       int    `json:"count"`
	CurrentPage int    `json:"currentPage"`
	TotalPage   int    `json:"totalPage"`
}

func NewPaginatedResponse[T any](resp PaginatedResponse[T], page, size int) *PaginatedResponse[T] {
	return &PaginatedResponse[T]{
		Message:     resp.Message,
		Result:      resp.Result,
		Count:       resp.Count,
		CurrentPage: (page - 1) * (size + 1),
		TotalPage:   int(math.Ceil(float64(resp.Count) / float64(size))),
	}
}

type SuccessResponse struct {
	Message string `json:"message"`
}

func NewSuccessResponse(message string) *SuccessResponse {
	return &SuccessResponse{
		Message: message,
	}
}

type ErrorResponse struct {
	Message    string    `json:"message"`
	StatusCode int       `json:"statusCode"`
	TimeStamp  time.Time `json:"timeStamp"`
}

func NewErrorResponse(message string, statusCode int) *ErrorResponse {
	return &ErrorResponse{
		Message:    message,
		StatusCode: statusCode,
		TimeStamp:  time.Now(),
	}
}
