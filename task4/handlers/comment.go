package handlers

import (
	"GoBlogProject/config"
	"GoBlogProject/models"
	"GoBlogProject/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 实现评论的创建功能，已认证的用户可以对文章发表评论。
// 实现评论的读取功能，支持获取某篇文章的所有评论列表。
func CommentCreateHandler(c *gin.Context) {
	id := c.Param("id")
	uid := c.GetInt("user_id")
	var comm models.Comment
	if err := c.ShouldBindJSON(&comm); err != nil {
		// c.JSON(http.StatusBadRequest, gin.H{"error": "数据绑定失败"})
		utils.ResponseError(c, http.StatusBadRequest, "数据绑定失败")
		return
	}
	pid, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("数据转换错误")
		return
	}
	comm.PostID = uint(pid)
	comm.UserID = uint(uid)
	if err := config.DB.Create(&comm).Error; err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"errro": "评论创建失败"})
		utils.ResponseError(c, http.StatusInternalServerError, "评论创建失败")
		return
	}
	var fullComment models.Comment
	if err := config.DB.Preload("User").Preload("Post").Preload("Post.User").First(&fullComment, comm.ID).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "查询评论失败")
		return
	}
	c.JSON(http.StatusOK, gin.H{"新增的评论是": fullComment})
}
func CommentGetHandler(c *gin.Context) {
	id := c.Param("id")
	pid, err := strconv.Atoi(id)
	if err != nil {
		// c.JSON(http.StatusBadRequest, gin.H{"error": "非法ID"})
		utils.ResponseError(c, http.StatusBadRequest, "非法ID")
		return
	}
	var comms []models.Comment
	if err := config.DB.Preload("User").Where("post_id=?", pid).Find(&comms).Error; err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"error": "获取评论失败"})
		utils.ResponseError(c, http.StatusInternalServerError, "获取评论失败")
		return
	}
	c.JSON(http.StatusOK, gin.H{"你想看到的文章的评论有：": comms})
}
