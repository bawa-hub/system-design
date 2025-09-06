package models

import (
	"time"

	"gorm.io/gorm"
)

// BaseModel contains common fields for all models
type BaseModel struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// APIResponse represents a standard API response
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// PaginationRequest represents pagination parameters
type PaginationRequest struct {
	Page     int `json:"page" form:"page" binding:"min=1"`
	PageSize int `json:"page_size" form:"page_size" binding:"min=1,max=100"`
}

// PaginationResponse represents paginated response
type PaginationResponse struct {
	Page       int         `json:"page"`
	PageSize   int         `json:"page_size"`
	Total      int64       `json:"total"`
	TotalPages int         `json:"total_pages"`
	Data       interface{} `json:"data"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Code    int    `json:"code,omitempty"`
}

// SuccessResponse creates a success response
func SuccessResponse(data interface{}, message ...string) APIResponse {
	msg := "Success"
	if len(message) > 0 && message[0] != "" {
		msg = message[0]
	}
	return APIResponse{
		Success: true,
		Message: msg,
		Data:    data,
	}
}

// ErrorResponse creates an error response
func NewErrorResponse(err string, code ...int) ErrorResponse {
	errorCode := 500
	if len(code) > 0 {
		errorCode = code[0]
	}
	return ErrorResponse{
		Success: false,
		Error:   err,
		Code:    errorCode,
	}
}

// CalculatePagination calculates pagination values
func CalculatePagination(page, pageSize int, total int64) PaginationResponse {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}

	totalPages := int((total + int64(pageSize) - 1) / int64(pageSize))
	if totalPages < 1 {
		totalPages = 1
	}

	return PaginationResponse{
		Page:       page,
		PageSize:   pageSize,
		Total:      total,
		TotalPages: totalPages,
	}
}
