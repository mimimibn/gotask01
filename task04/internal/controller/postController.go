package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"task04/internal/common"
	"task04/internal/common/response"
	"task04/internal/model"
	"task04/internal/repository"
	"time"
)

func CreatePost(c *gin.Context) {
	var postAddDTO model.PostAddDTO
	if err := c.ShouldBindJSON(&postAddDTO); err != nil {
		response.FailMsg("参数错误", c)
	}
	session := common.GetCurrentUserSession(c)
	post := model.Post{
		UserId:  session.Id,
		Title:   postAddDTO.Title,
		Content: postAddDTO.Content,
		BaseEntity: model.BaseEntity{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	if err := repository.DB.Debug().Create(&post).Error; err != nil {
		response.FailMsg("保存数据失败", c)
		return
	}
	fmt.Println(postAddDTO)
}

/*
*
这里少加了分页，
*/
func FindListPost(c *gin.Context) {
	session := common.GetCurrentUserSession(c)
	postArray := []model.Post{}
	repository.DB.Debug().Model(&model.Post{}).Where("user_id = ?", session.Id).Find(&postArray)
	response.Data(postArray, c)
}
func FindByIdPost(c *gin.Context) {
	id := c.Query("id")
	session := common.GetCurrentUserSession(c)
	post := model.Post{}
	repository.DB.Debug().Model(&model.Post{}).Where("user_id = ? AND id = ?", session.Id, id).First(&post)
	response.Data(post, c)
}
func UpdateByIdPost(c *gin.Context) {
	var postUpdDTO model.PostUpdDTO
	if err := c.ShouldBindJSON(&postUpdDTO); err != nil {
		response.FailMsg("参数错误", c)
	}
	session := common.GetCurrentUserSession(c)
	post := model.Post{}
	err := repository.DB.Debug().Model(&model.Post{}).Where("id = ?", postUpdDTO.Id).First(&post).Error
	if err != nil {
		response.FailMsg("数据不存在", c)
	}
	if post.UserId != session.Id {
		response.FailMsg("无权限", c)
	}
	post.UpdatedAt = time.Now()
	post.Title = postUpdDTO.Title
	post.Content = postUpdDTO.Content
	repository.DB.Debug().Model(&post).Updates(post)
	response.OK(c)
}
func DeleteByIdPost(c *gin.Context) {
	id := c.Param("id")
	session := common.GetCurrentUserSession(c)
	post := model.Post{}
	err := repository.DB.Debug().Model(&model.Post{}).Where("id = ?", id).First(&post).Error
	if err != nil {
		response.FailMsg("数据不存在", c)
	}
	if post.UserId != session.Id {
		response.FailMsg("无权限", c)
	}
	repository.DB.Debug().Model(&model.Post{}).Delete(&post)
	response.OK(c)
}
