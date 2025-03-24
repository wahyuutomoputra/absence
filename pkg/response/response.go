package response

import "github.com/gin-gonic/gin"

// Response represents the standard API response structure
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Success sends a success response
func Success(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, Response{
		Status:  true,
		Message: message,
		Data:    data,
	})
}

// Error sends an error response
func Error(c *gin.Context, code int, message string) {
	c.JSON(code, Response{
		Status:  false,
		Message: message,
	})
}
