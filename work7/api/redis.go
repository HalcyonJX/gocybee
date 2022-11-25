package api

import (
	"github.com/gin-gonic/gin"
	"gocybee/work7/dao"
	"gocybee/work7/utils"
)

func LikeSet(c *gin.Context) {
	username := c.PostForm("username")
	if username == "" {
		utils.RespFail(c, "请输入用户名")
	}
	flag := dao.IsLike(username)
	if flag {
		dao.LikeDel(username)
		utils.RespSuccess(c, "删除点赞成功")
	} else {
		dao.LikeSet(username)
		utils.RespSuccess(c, "点赞成功")
	}
}
