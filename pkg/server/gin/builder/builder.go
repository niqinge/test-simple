package builder

import "github.com/niqinge/test-simple/pkg/server/gin"

func NewHandler() *gin.Handler{
	return &gin.Handler{
		Manager: nil,
	}
}
