package controller

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"task04/internal/common/response"
	"task04/internal/model"
	"task04/internal/repository"
	"time"
)

func Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.FailMsg("参数错误", c)
		//c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		response.FailMsg("密码加密失败", c)
		//c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)
	user.BaseEntity.CreatedAt = time.Now()
	user.BaseEntity.UpdatedAt = time.Now()
	if err := repository.DB.Create(&user).Error; err != nil {
		response.FailMsg("注册失败", c)
		//c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	response.OK(c)
	//c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.FailMsg("参数错误", c)
		//c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var storedUser model.User
	if err := repository.DB.Where("user_name = ?", user.UserName).First(&storedUser).Error; err != nil {
		response.FailMsg("用户不存在", c)
		//c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		response.FailMsg("用户密码校验失败", c)
		//c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// 生成 JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       storedUser.Id,
		"username": storedUser.UserName,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	//这里key可以取自配置文件
	tokenString, err := token.SignedString([]byte("123456"))
	if err != nil {
		response.FailMsg("成功token失败", c)
		//c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	//直接返回
	mapData := make(map[string]string)
	mapData["token"] = tokenString
	response.Data(mapData, c)
	//c.JSON(http.StatusOK, tokenString)
}
