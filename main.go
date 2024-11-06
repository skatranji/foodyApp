package main

import (
	"myapp/db"
	"myapp/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Database
	db.Init()

	router := gin.Default()

	// Authentication Routes
	router.POST("/api/auth/signup", handlers.SignUpHandler)
	router.POST("/api/auth/login", handlers.LoginHandler)
	router.GET("/api/auth/refresh", handlers.RefreshTokenHandler)

	// User Routes
	router.GET("/api/user/profile/:userId", handlers.GetUserProfileHandler)
	router.PUT("/api/user/profile", handlers.UpdateUserProfileHandler)
	router.GET("/api/user/following", handlers.GetUserFollowingHandler)
	router.POST("/api/user/follow", handlers.FollowUserHandler)

	// Video Routes
	router.POST("/api/videos/upload", handlers.UploadVideoHandler)
	router.GET("/api/videos/feed", handlers.GetFeedHandler)
	router.GET("/api/videos/:videoId", handlers.GetVideoDetailsHandler)
	router.DELETE("/api/videos/:videoId", handlers.DeleteVideoHandler)
	router.POST("/api/videos/like", handlers.LikeVideoHandler)
	router.GET("/api/videos/likes/:videoId", handlers.GetVideoLikesHandler)
	router.POST("/api/videos/comment", handlers.AddCommentHandler)
	router.GET("/api/videos/comments/:videoId", handlers.GetCommentsHandler)
	router.DELETE("/api/videos/comments/:commentId", handlers.DeleteCommentHandler)

	// Search Routes
	router.GET("/api/search/videos", handlers.SearchVideosHandler)
	router.GET("/api/search/users", handlers.SearchUsersHandler)

	// Notification Routes
	router.GET("/api/notifications", handlers.GetNotificationsHandler)
	router.POST("/api/notifications/mark-as-read", handlers.MarkNotificationsAsReadHandler)

	// Admin Routes
	router.GET("/api/admin/videos", handlers.GetAllVideosHandler)
	router.PUT("/api/admin/videos/:videoId", handlers.UpdateVideoDetailsHandler)
	router.DELETE("/api/admin/users/:userId", handlers.DeleteUserHandler)

	router.Run(":8080")
}
