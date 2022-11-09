package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yunus-mohammed/golang-project/models"
)

func RespController(c *gin.Context) {

	var req models.Details

	statusCode := http.StatusOK

	err := c.ShouldBind(&req)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}

	res, err := models.GetData(req, c)

	if err != nil {
		statusCode = http.StatusInternalServerError
	}

	c.JSON(statusCode, res)

	return
}
