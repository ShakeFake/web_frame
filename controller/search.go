package controller

import (
	log "github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
	"wilikidi/socket/model"
	"wilikidi/socket/service"
	"wilikidi/socket/utils"
)

func RrsTaskList(c *gin.Context) {
	var s model.Student
	if err := c.BindJSON(&s); err != nil {
		utils.GenerateReturnData(c, utils.ParamsError, err.Error(), "")
		return
	}

	log.Infof("request parameters is: %v", s)

	response := service.InsertAStudent(s)

	utils.GenerateReturnData(c, utils.SUCCESS, "", response)

}

func Version(c *gin.Context) {
	utils.GenerateReturnData(c, utils.SUCCESS, "success", "1.0")
}
