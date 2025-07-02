package utils

import "github.com/gin-gonic/gin"

// Format respons sukses
func SuccessResponse(data interface{}) gin.H {
    return gin.H{
        "success": true,
        "data":    data,
    }
}

// Format respons error
func ErrorResponse(message string) gin.H {
    return gin.H{
        "success": false,
        "error":   message,
    }
}
