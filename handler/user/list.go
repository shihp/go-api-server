package user

import (
	"apiserver/handler"
	"apiserver/pkg/errno"
	"apiserver/service"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"strconv"
)

// List list the users in the database.
func List(c *gin.Context) {
	log.Info("list function called")
	var r ListRequest
	username := c.Query("username")
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	infos, count, err := service.ListUser(username, offset, limit)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}

	handler.SendResponse(c, nil, ListResponse{
		TotalCount: count,
		UserList:   infos,
	})
}
