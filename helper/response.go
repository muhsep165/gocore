package helper

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ResponseList struct {
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
	Page      int         `json:"page"`
	TotalPage int         `json:"total_page"`
	Limit     int         `json:"limit"`
	OrderBy   []string    `json:"order_by"`
}

func ERR_NOT_FOUND() Response {
	return Response{
		Success: false,
		Message: "Data Invalid or Not Found",
		Data:    nil,
	}
}

func ERR_BAD_REQUEST() Response {
	return Response{
		Success: false,
		Message: "Bad Request, Validation Error",
		Data:    nil,
	}
}

func ERR_UNAUTHORIZED() Response {
	return Response{
		Success: false,
		Message: "Unauthorized",
		Data:    nil,
	}
}

func ERR_SERVER_ERROR() Response {
	return Response{
		Success: false,
		Message: "Service Unavailable",
		Data:    nil,
	}
}
