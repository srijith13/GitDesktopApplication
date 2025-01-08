package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"git-visualizer/app/dto"
	"git-visualizer/app/helper"
	"git-visualizer/app/service"
)

func Ping(c *gin.Context) {
	fmt.Println("Pimg")
	c.JSON(http.StatusOK, helper.BuildResponse("Pong", "done "))
}

func GitCloneRepo(c *gin.Context) {
	var request dto.Request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, helper.BuildErrorResponse("Bad Request", err.Error(), "Failed to clone repo"))
		return
	}
	// if errors := helper.ValidateCustomBody(request); len(errors) != 0 {
	// 	c.JSON(http.StatusBadRequest, helper.BuildErrorResponse("Bad Request", errors, nil, false))
	// 	return
	// }

	// connection of database
	// code := config.GetCode(c.GetHeader("city"))
	// dbCon := config.CreateDbConPool(code)
	// request.Code = code
	// request.Limit = 4

	// connection of database
	// go helper.AddSearch(c, &request, "GetBusRoute")
	// var busRes []dto.RouteResponseV2
	// var moreData bool
	// busResChan := make(chan []dto.RouteResponseV2)
	// moreDataChan := make(chan bool)
	err := service.GitCloneRepo(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, helper.BuildErrorResponse("Bad Request", err.Error(), "Failed to clone repo"))
		return
	}

	// go service.GetBusRoute(dbCon, &request, busResChan, moreDataChan)
	// busRes = <-busResChan
	// moreData = <-moreDataChan

	c.JSON(http.StatusOK, helper.BuildResponse("Repository Clone ", "Success"))
}

func GetRepoBranches(c *gin.Context) {
	var request dto.Request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, helper.BuildErrorResponse("Bad Request", err.Error(), "Failed to get branches"))
		return
	}
	branchLists, err := service.GetRepoBranches(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.BuildErrorResponse("Bad Request", err.Error(), "Failed to get branches"))
		return
	}
	c.JSON(http.StatusOK, helper.BuildResponse("Repository Branches", branchLists))
}

func GetBranchStatus(c *gin.Context) {
	var request dto.Request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, helper.BuildErrorResponse("Bad Request", err.Error(), "Failed to get status"))
		return
	}
	branchStatus, err := service.GetBranchStatus(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.BuildErrorResponse("Bad Request", err.Error(), "Failed to get status"))
		return
	}
	c.JSON(http.StatusOK, helper.BuildResponse("Repository Branch Status", branchStatus))
}

func GetLogs(c *gin.Context) {
	var request dto.Request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, helper.BuildErrorResponse("Bad Request", err.Error(), "Failed to get logs"))
		return
	}
	commitLogs, err := service.GetLogs(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.BuildErrorResponse("Bad Request", err.Error(), "Failed to get logs"))
		return
	}
	c.JSON(http.StatusOK, helper.BuildResponse("Repository Branch Status", commitLogs))
}

func GitCheckoutBranches(c *gin.Context) {
	var request dto.Request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, helper.BuildErrorResponse("Bad Request", err.Error(), fmt.Sprintf("Failed checked out to branch %s ", request.Branch)))
		return
	}
	err := service.GitCheckoutBranches(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.BuildErrorResponse("Bad Request", err.Error(), fmt.Sprintf("Failed checked out to branch %s ", request.Branch)))
		return
	}
	c.JSON(http.StatusOK, helper.BuildResponse("Checkout Branch ", fmt.Sprintf("Successfully checked out to branch %s ", request.Branch)))
}
