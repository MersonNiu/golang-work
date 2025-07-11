package handlers

import (
	"GoBlogProject/config"
	"GoBlogProject/models"
	"GoBlogProject/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(c *gin.Context) {
	//读取参数
	var req models.User
	if err := c.ShouldBindJSON(&req); err != nil {
		// c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误"})
		utils.ResponseError(c, http.StatusBadRequest, "请求参数错误")
		return
	}
	//判断用户是否已存在
	var existing models.User
	if err := config.DB.Where("Username=?", req.Username).First(&existing).Error; err == nil {
		// c.JSON(http.StatusBadRequest, gin.H{"error": "用户名已存在"})
		utils.ResponseError(c, http.StatusBadRequest, "用户名已存在")
		return
	}
	//密码加密
	hashed, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	req.Password = string(hashed)
	//写入数据库
	if err := config.DB.Create(&req).Error; err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"err": "创建用户失败，数据库数据写入失败"})
		utils.ResponseError(c, http.StatusInternalServerError, "创建用户失败，数据库数据写入失败")
		return
	}
	c.JSON(200, gin.H{"id": req.ID, "username": req.Username})
}

func LoginHandler(c *gin.Context) {
	//读取数据
	var req models.User
	if err := c.ShouldBindJSON(&req); err != nil {
		// c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误"})
		utils.ResponseError(c, http.StatusBadRequest, "请求参数错误")
		return
	}
	//查看用户是否存在
	var user models.User
	if err := config.DB.Where("Username=?", req.Username).First(&user).Error; err != nil {
		// c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名错误"})
		utils.ResponseError(c, http.StatusUnauthorized, "用户名错误")
		return
	}
	//不存在则返回提示，存在则比较密码
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		// c.JSON(http.StatusUnauthorized, gin.H{"error": "密码错误"})
		utils.ResponseError(c, http.StatusUnauthorized, "密码错误")
		return
	}
	token := utils.GenerateToken(int(user.ID))
	c.JSON(http.StatusAccepted, gin.H{"token": token})
}

func MeHandler(c *gin.Context) {
	uid := c.GetInt("user_id")
	var user models.User
	if err := config.DB.First(&user, uid).Error; err != nil {
		// c.JSON(404, gin.H{"error": "用户不存在"})
		utils.ResponseError(c, http.StatusNotFound, "用户不存在")
		return
	}
	c.JSON(200, gin.H{"id": user.ID, "username": user.Username})
}
