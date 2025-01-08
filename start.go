package main

import (
	// "fmt"
	"git-visualizer/app"
	"git-visualizer/app/config"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	// "os"
	// "github.com/gin-gonic/gin"
	// "honnef.co/go/tools/config"
	// "github.com/yourusername/github-gitlab-profile-app/api/router"
	// "github.com/yourusername/github-gitlab-profile-app/config"
	// "github.com/yourusername/github-gitlab-profile-app/internal/core/adapter"
	// "github.com/yourusername/github-gitlab-profile-app/internal/core/service"
)

// func main() {
// 	// Load configuration from environment variables or config files
// 	// Load configuration from environment variables or config files
// 	config.Load()

// 	// Initialize GitHub and GitLab adapters with access tokens
// 	githubAdapter := github.NewGitHubAdapter(os.Getenv("GITHUB_TOKEN"))
// 	// gitlabAdapter := adapter.NewGitLabAdapter(os.Getenv("GITLAB_TOKEN"))

// 	// Create service instances
// 	githubService := service.GitHubProfileService(githubAdapter)
// 	// gitlabService := service.NewGitLabProfileService(gitlabAdapter)

// 	// GitAdapter (using the current directory as an example)
// 	gitAdapter := adapter.NewGitAdapter("/path/to/your/repo")
// 	gitService := service.NewGitService(gitAdapter)
// 	gitHandler := handlers.NewGitHandler(gitService)

// 	// Set up Gin router and API handlers
// 	r := gin.Default()

// 	// Create ProfileHandler with injected services
// 	handler := handlers.NewProfileHandler(githubService, gitlabService)

// 	// Register routes with Gin
// 	r.GET("/profile/github", handler.GetGitHubProfile)
// 	r.GET("/profile/gitlab", handler.GetGitLabProfile)

// 	// Start the Gin server
// 	port := ":8080"
// 	fmt.Printf("Server is running on %s...\n", port)
// 	if err := r.Run(port); err != nil {
// 		log.Fatalf("Failed to start server: %v", err)
// 	}

// }

func init() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
}

func main() {
	loc, _ := time.LoadLocation(config.GinTZ)
	time.Local = loc
	gin.SetMode(config.GinMode)

	log.Println("Your Application Started")
	start := app.App("Git Visualizer App")
	start.StartApp()
}
