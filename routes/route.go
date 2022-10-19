package routes

import (
	"MyGramHacktiv8/database"
	"MyGramHacktiv8/handler"
	"MyGramHacktiv8/middleware"
	"MyGramHacktiv8/pkg/helpers"
	"MyGramHacktiv8/repository/commentRepository/comment_pg"
	"MyGramHacktiv8/repository/photoRepository/photo_pg"
	"MyGramHacktiv8/repository/socialmediaRepository/socialmedia_pg"
	"MyGramHacktiv8/repository/userRepository/user_pg"
	"MyGramHacktiv8/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartApp() {
	database.StartDB()
	db := database.GetDB()

	userRepo := user_pg.NewUserPG(db)
	userService := service.NewUserService(userRepo)
	userRestHandler := handler.NewUserRestHandler(userService)

	photoRepo := photo_pg.NewPhotoPG(db)
	photoService := service.NewPhotoService(photoRepo)
	photoRestHandler := handler.NewPhotoRestHandler(photoService)

	commentRepo := comment_pg.NewCommentPG(db)
	commentService := service.NewCommentService(commentRepo)
	commentRestHandler := handler.NewCommentRestHandler(commentService)

	socialMediaRepo := socialmedia_pg.NewSocialMediaPG(db)
	socialMediaService := service.NewSocialMediaService(socialMediaRepo)
	socialMediaRestHandler := handler.NewSocialMediaRestHandler(socialMediaService)

	authMiddleware := middleware.NewAuthService(userRepo, photoRepo, commentRepo, socialMediaRepo)

	// ! Routing
	router := gin.Default()
	router.Use(cors.Default())
	v1 := router.Group("/api/v1")

	v1.POST("/login", userRestHandler.Login)
	v1.POST("/register", userRestHandler.Register)
	userRoute := v1.Group("/users")
	userRoute.PUT("/:userID", authMiddleware.Authentication(), userRestHandler.UpdateUserData)
	userRoute.DELETE("/", authMiddleware.Authentication(), userRestHandler.DeleteUser)

	photoRoute := v1.Group("/photos")
	photoRoute.Use(authMiddleware.Authentication())
	photoRoute.POST("/", photoRestHandler.PostPhoto)
	photoRoute.GET("/", photoRestHandler.GetAllPhotos)
	photoRoute.PUT("/:photoID", authMiddleware.PhotoAuthorization(), photoRestHandler.UpdatePhoto)
	photoRoute.DELETE("/:photoID", authMiddleware.PhotoAuthorization(), photoRestHandler.DeletePhoto)

	commentRoute := v1.Group("/comments")
	commentRoute.Use(authMiddleware.Authentication())
	commentRoute.POST("/", commentRestHandler.PostComment)
	commentRoute.GET("/", commentRestHandler.GetAllComments)
	commentRoute.PUT("/:commentID", authMiddleware.CommentAuthorization(), commentRestHandler.UpdateComment)
	commentRoute.DELETE("/:commentID", authMiddleware.CommentAuthorization(), commentRestHandler.DeleteComment)

	socialMediaRoute := v1.Group("/social-medias")
	socialMediaRoute.Use(authMiddleware.Authentication())
	socialMediaRoute.POST("/", socialMediaRestHandler.AddSocialMedia)
	socialMediaRoute.GET("/", socialMediaRestHandler.GetAllSocialMedias)
	socialMediaRoute.PUT("/:socialMediaID", authMiddleware.SocialMediaAuthorization(), socialMediaRestHandler.EditSocialMediaData)
	socialMediaRoute.DELETE("/:socialMediaID", authMiddleware.SocialMediaAuthorization(), socialMediaRestHandler.DeleteSocialMedia)

	router.Run(":" + helpers.GodotEnv("APP_PORT"))
}
