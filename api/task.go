package api

import (
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
	"todo/pkg/utils"
	"todo/service"
)

func CreateTask(c *gin.Context) {
	var createTask service.CreateTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createTask); err != nil {
		res := createTask.Create(claim.Id)
		c.JSON(200, res)
	} else {
		logging.Error()
		c.JSON(400, err)
	}
}
