package models

type (
	BaseResponse struct {
		// Success Успешно ли выполнен запрос
		Success bool `json:"success"`
		// Error Сообщение об ошибке
		Error string `json:"error,omitempty"`
		// ErrorCode Код ошибки
		ErrorCode int `json:"errorCode"`
	}
)
