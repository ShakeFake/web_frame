package controller

import (
	log "github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
	"wilikidi/gin/model"
	"wilikidi/gin/service"
	"wilikidi/gin/utils"
)

func Validator(c *gin.Context) {
	var s model.Student
	if err := c.BindJSON(&s); err != nil {
		utils.GenerateReturnData(c, utils.ParamsError, utils.ErrorTranslate(err), "")
		return
	}

	log.Infof("request parameters is: %v", s)

	response := service.InsertAStudent(s)
	_ = response

	utils.GenerateReturnData(c, utils.SUCCESS, "", s)

}

func Version(c *gin.Context) {
	utils.GenerateReturnData(c, utils.SUCCESS, "success", "1.0")
}
