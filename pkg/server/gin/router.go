package gin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	s       *http.Server
	handler *Handler
}

func (s *Server) GetEngine() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	e := gin.New()

	r := e.Group("api/v1/")
	// r.Use() // 校验token等
	{
		// 获取单一用户详情
		r.POST("item/detail", s.handler.QueryUser())
		// 修改用户信息
		r.POST("item/update", s.handler.UpdateUser())
		// 删除用户
		r.POST("item/del", s.handler.DelUser())
		// 新增用户
		r.POST("item/add", s.handler.AddUser())
		// 查询用户列表
		r.POST("item/list", s.handler.QueryUserList())
	}

	return e
}
