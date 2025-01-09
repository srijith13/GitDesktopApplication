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
	// General details
	visualizerApp.POST("/cloneRepo", controller.GitCloneRepo)
	visualizerApp.POST("/getRepoBranches", controller.GetRepoBranches)
	visualizerApp.POST("/getBranchStatus", controller.GetBranchStatus)
	visualizerApp.POST("/getLog", controller.GetLogs)
	visualizerApp.POST("/checkoutBranch", controller.GitCheckoutBranches)
	visualizerApp.POST("/stashChanges", controller.GitStash)

	// push pull commit details
	visualizerApp.POST("/addCommitFiles", controller.GitAddCommitFiles)
	visualizerApp.POST("/commitChanges", controller.GitCommitChanges)
	visualizerApp.POST("/deleteLastCommit", controller.GitDeleteLastCommit)
	visualizerApp.POST("/pushChanges", controller.GitPushChanges)
	visualizerApp.POST("/pullChanges", controller.GitPullChanges)
	// /merge need to add.
	// Merge Conflict Resolution
	// POST /merge/resolve: This endpoint resolves merge conflicts using the strategy provided (theirs, ours, etc.). It attempts to merge with the specified strategy and then commits the result.
	// GET /merge/conflicts: This endpoint checks if there are unresolved merge conflicts in the repository. If there are files in conflict, it returns a message indicating so.
	// Rebase
	// POST /rebase: This endpoint performs a rebase operation. It rebases the current branch onto the provided target branch.

	// Visualizer with blame details
	visualizerApp.POST("/blame", controller.GitBlame)
	visualizerApp.POST("/branchVisualize", controller.GitBranchVisualize)
	visualizerApp.POST("/incomingChanges", controller.GitIncomingChanges)

	// // other git commands
	// Cherry-pick
	// POST /cherry-pick: Cherry-picks a commit from one branch and applies it to the current branch. This can be useful to select specific commits without merging entire branches.
	// Revision
	// GET /commit/revision: Fetches the revision details of a commit given its hash.
	// Reorder Commits
	// POST /reorder-commits: Reorders commits interactively using a rebase.
	// Squash/Fixup
	// POST /squash-commits: Squashes a series of commits into one. It retains commit messages unless the fixup flag is used.
	// POST /fixup-commits: Fixup commits by squashing them into the previous commit, discarding the commit message of the fixed-up commits.

	// hub := visualizerApp.Group("hub")
	// lab := visualizerApp.Group("lab")

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
