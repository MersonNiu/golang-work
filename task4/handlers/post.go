package handlers

import (
	"GoBlogProject/config"
	"GoBlogProject/models"
	"GoBlogProject/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 实现文章的创建功能，只有已认证的用户才能创建文章，创建文章时需要提供文章的标题和内容。

func PostCreateHandler(c *gin.Context) {
	var req struct {
		Title   string
		Content string
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		// c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误"})
		utils.ResponseError(c, http.StatusBadRequest, "请求参数错误")
		return
	}
	uid := c.GetInt("user_id")
	post := models.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  uint(uid),
	}
	if err := config.DB.Create(&post).Error; err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"error": "创建文章失败"})
		utils.ResponseError(c, http.StatusInternalServerError, "创建文章失败")
	}
	var fullPost models.Post
	if err := config.DB.Preload("User").First(&fullPost, post.ID).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "查询文章失败")
		return
	}
	c.JSON(http.StatusCreated, fullPost)
}

// 实现文章的读取功能，支持获取所有文章列表和单个文章的详细信息。
func PostGetAllHandler(c *gin.Context) {
	var post []models.Post
	if err := config.DB.Preload("User").Find(&post).Error; err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"error": "获取文章列表失败"})
		utils.ResponseError(c, http.StatusInternalServerError, "获取文章列表失败")
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"post": post})
}
func PostGetIdHandler(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	if err := config.DB.Preload("User").First(&post, id).Error; err != nil {
		// c.JSON(http.StatusNotFound, gin.H{"error": "获取文章失败"})
		utils.ResponseError(c, http.StatusNotFound, "获取文章失败")
		return
	}
	c.JSON(http.StatusOK, gin.H{"想查看的文章是：": post})
}

// 实现文章的更新功能，只有文章的作者才能更新自己的文章。
func PostUpdateHandler(c *gin.Context) {
	//先获取要操作的文章的所有信息
	var post models.Post
	id := c.Param("id")

	if err := config.DB.First(&post, id).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "文章查询失败")
		return
	}
	uid := c.GetInt("user_id")
	if post.UserID != uint(uid) {
		utils.ResponseError(c, http.StatusForbidden, "无权限修改")
		return
	}
	var seq models.Post
	if err := c.ShouldBindJSON(&seq); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, "数据绑定失败")
		return
	}
	post.Title = seq.Title
	post.Content = seq.Content
	if err := config.DB.Save(&post).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "更新失败")
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "文章更新成功", "post": post})

}

// 实现文章的删除功能，只有文章的作者才能删除自己的文章。
func PostDeleteHandler(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	if err := config.DB.First(&post, id).Error; err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"error": "数据查询失败"})
		utils.ResponseError(c, http.StatusInternalServerError, "数据查询失败")
		return
	}
	uid := c.GetInt("user_id")
	if post.UserID != uint(uid) {
		// c.JSON(http.StatusForbidden, gin.H{"error": "无权限删除"})
		utils.ResponseError(c, http.StatusForbidden, "无权限删除")
		return
	}
	if err := config.DB.Delete(&post).Error; err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"error": "删除失败"})
		utils.ResponseError(c, http.StatusInternalServerError, "删除失败")
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "文章删除成功"})
}
