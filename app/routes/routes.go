package routes

import (
	"fmt"
	"git-visualizer/app/config"
	"git-visualizer/app/controller"

	"github.com/gin-gonic/gin"
)

func Routes() {
	router := gin.Default()
	fmt.Println("its here", config.AppPort)
	// each api request call 15 in one minute according to UserId
	// for a given UserId has a maximum limit of 1000 request in a minute
	// rateLimiter := limiter.NewRateLimiter(time.Minute, 1000, func(ctx *gin.Context) (string, error) {
	// 	key := ctx.Request.Header.Get("UserId")
	// 	if key != "" {
	// 		return key, nil
	// 	}
	// 	return "", nil
	// })

	// router.Use(auth.RequestID())

	// router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
	// 	// your custom format
	// 	return fmt.Sprintf("%s | %s | %s | %s | ",
	// 		param.Keys["request-id"],
	// 		param.Request.Header.Get("userId"),
	// 		param.Request.Header.Get("userName")
	// }))

	// router.Use(rateLimiter.Middleware())
	router.GET("/ping", controller.Ping)

	visualizerApp := router.Group("visualizer/v1")

	// visualizerApp.Use(middleware.AppAutherization)
	visualizerApp.POST("/cloneRepo", controller.GitCloneRepo)
	visualizerApp.POST("/getRepoBranches", controller.GetRepoBranches)
	visualizerApp.POST("/getBranchStatus", controller.GetBranchStatus)
	visualizerApp.POST("/getLog", controller.GetLogs)
	visualizerApp.POST("/checkoutBranch", controller.GitCheckoutBranches)

	// visualizerApp.POST("/getNearByMetroStation", controller.GetNearByMetroStation)
	// visualizerApp.POST("/getNearByBusStop", controller.GetNearByBusStop)
	// visualizerApp.POST("/searchMetroStation", controller.SearchMetroStation)
	// visualizerApp.POST("/searchBusStop", controller.SearchBusStop)
	// visualizerApp.POST("/getAllMetroRoute", controller.GetAllMetroRoute)
	// visualizerApp.POST("/getMetroScheduleByRouteId", controller.GetMetroScheduleByRouteId)
	// visualizerApp.POST("/getNearByBusRoute", controller.GetNearByBusRoute)
	// visualizerApp.POST("/searchBusRoute", controller.SearchBusRoute)
	// visualizerApp.GET("/getBusTicketingRoute", controller.GetBusTicketingRoute)
	// visualizerApp.Use(middleware.AuthRequired) // when auth come need to add the middleware

	router.Run(fmt.Sprintf(`:%s`, config.AppPort))
}
