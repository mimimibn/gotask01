package controller

import (
	"github.com/gin-gonic/gin"
	"task04/internal/common"
	"task04/internal/common/response"
	"task04/internal/model"
	"task04/internal/repository"
	"time"
)

func CreateComment(c *gin.Context) {
	var commentAddDTO model.CommentAddDTO
	if err := c.ShouldBindJSON(&commentAddDTO); err != nil {
		response.FailMsg("参数错误", c)
		return
	}
	session := common.GetCurrentUserSession(c)
	comment := model.Comment{
		UserId:  session.Id,
		PostId:  commentAddDTO.PostId,
		Content: commentAddDTO.Content,
		BaseEntity: model.BaseEntity{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	repository.DB.Debug().Save(&comment)
	response.OK(c)
}
func FindCommentListByPostId(c *gin.Context) {
	id := c.Query("id")
	comments := []model.Comment{}
	repository.DB.Debug().Model(&model.Comment{}).Where("post_id = ?", id).Find(&comments)
	response.Data(comments, c)
}
