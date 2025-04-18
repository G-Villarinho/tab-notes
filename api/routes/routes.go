package routes

import (
	"database/sql"
	"strings"

	"github.com/g-villarinho/tab-notes-api/clients"
	"github.com/g-villarinho/tab-notes-api/configs"
	"github.com/g-villarinho/tab-notes-api/handlers"
	"github.com/g-villarinho/tab-notes-api/middlewares"
	"github.com/g-villarinho/tab-notes-api/pkgs"
	"github.com/g-villarinho/tab-notes-api/repositories"
	"github.com/g-villarinho/tab-notes-api/services"
)

func SetupRoutes(db *sql.DB, queueClient clients.QueueClient) *Router {
	router := NewRouter()

	if strings.ToLower(configs.Env.Env) == "development" {
		envHandler := handlers.NewEnvironmentHandler()

		router.GET("/envs", envHandler.GetEnvs)
	}

	setupHealthRoutes(db, router)
	setupAuthRoutes(db, router, queueClient)
	setupRegisterRoutes(db, router)
	setupUserRoutes(db, router)
	setupFollowerRoutes(db, router)
	setupSessionRoutes(db, router)
	setupPostRoutes(db, router)
	setupFeedRoutes(db, router)

	return router
}

func setupHealthRoutes(db *sql.DB, router *Router) {
	healthService := services.NewHealthService(db)
	healthHandler := handlers.NewHealthHandler(healthService)

	router.GET("/health", healthHandler.Check)
}

func setupAuthRoutes(db *sql.DB, router *Router, queueClient clients.QueueClient) {
	ecdsa := pkgs.NewEcdsaKeyPair()
	requestContext := pkgs.NewRequestContext()

	queueService := services.NewQueueService(queueClient)

	tokenService := services.NewTokenService(ecdsa)
	sessionRepository := repositories.NewSessionRepository(db)

	userRepository := repositories.NewUserRepository(db)
	followerRepository := repositories.NewFollowerRepository(db)
	followerService := services.NewFollowerService(followerRepository, userRepository)
	userService := services.NewUserService(followerService, userRepository)

	sessionService := services.NewSessionService(tokenService, sessionRepository)
	authService := services.NewAuthService(sessionService, userService, queueService)
	authHandler := handlers.NewAuthHandler(authService, requestContext)

	authMiddleware := middlewares.NewAuthMiddleware(ecdsa, requestContext, sessionService)

	router.POST("/authenticate", authHandler.SendAuthenticationLink)
	router.GET("/magic-link/authenticate", authHandler.AuthenticateFromLink)
	router.POST("/logout", authMiddleware.Authenticated(authHandler.Logout))
}

func setupRegisterRoutes(db *sql.DB, router *Router) {
	ecdsa := pkgs.NewEcdsaKeyPair()
	tokenService := services.NewTokenService(ecdsa)

	sessionRepository := repositories.NewSessionRepository(db)
	userRepository := repositories.NewUserRepository(db)
	followerRepository := repositories.NewFollowerRepository(db)
	followerService := services.NewFollowerService(followerRepository, userRepository)
	userService := services.NewUserService(followerService, userRepository)

	sessionService := services.NewSessionService(tokenService, sessionRepository)
	registerService := services.NewRegisterService(userService, sessionService)
	registerHandler := handlers.NewRegisterHandler(registerService)

	router.POST("/register", registerHandler.RegisterUser)
}

func setupUserRoutes(db *sql.DB, router *Router) {
	ecdsa := pkgs.NewEcdsaKeyPair()
	requestContext := pkgs.NewRequestContext()

	tokenService := services.NewTokenService(ecdsa)
	sessionRepository := repositories.NewSessionRepository(db)
	sessionService := services.NewSessionService(tokenService, sessionRepository)

	authMiddleware := middlewares.NewAuthMiddleware(ecdsa, requestContext, sessionService)

	userRepository := repositories.NewUserRepository(db)
	followerRepository := repositories.NewFollowerRepository(db)
	followerService := services.NewFollowerService(followerRepository, userRepository)

	userService := services.NewUserService(followerService, userRepository)
	userHandler := handlers.NewUserHandler(requestContext, userService)

	router.GET("/me", authMiddleware.Authenticated(userHandler.GetProfile))
	router.GET("/users", authMiddleware.Authenticated(userHandler.SearchUsers))
	router.GET("/users/{username}", authMiddleware.Authenticated(userHandler.GetProfileByUsername))
	router.PUT("/users", authMiddleware.Authenticated(userHandler.UpdateUser))
}

func setupFollowerRoutes(db *sql.DB, router *Router) {
	ecdsa := pkgs.NewEcdsaKeyPair()
	requestContext := pkgs.NewRequestContext()

	tokenService := services.NewTokenService(ecdsa)
	sessionRepository := repositories.NewSessionRepository(db)
	sessionService := services.NewSessionService(tokenService, sessionRepository)

	userRepository := repositories.NewUserRepository(db)
	followerRepository := repositories.NewFollowerRepository(db)
	followerService := services.NewFollowerService(followerRepository, userRepository)
	followerHandler := handlers.NewFollowerHandler(requestContext, followerService)

	authMiddleware := middlewares.NewAuthMiddleware(ecdsa, requestContext, sessionService)

	router.POST("/users/{username}/follow", authMiddleware.Authenticated(followerHandler.FollowUser))
	router.POST("/users/{username}/unfollow", authMiddleware.Authenticated(followerHandler.UnfollowUser))
	router.GET("/users/{username}/followers", authMiddleware.Authenticated(followerHandler.GetFollowers))
	router.GET("/users/{username}/following", authMiddleware.Authenticated(followerHandler.GetFollowing))
	router.GET("/me/followers", authMiddleware.Authenticated(followerHandler.GetMyFollowers))
	router.GET("/me/following", authMiddleware.Authenticated(followerHandler.GetMyFollowing))
}

func setupSessionRoutes(db *sql.DB, router *Router) {
	ecdsa := pkgs.NewEcdsaKeyPair()
	requestContext := pkgs.NewRequestContext()

	tokenService := services.NewTokenService(ecdsa)
	sessionRepository := repositories.NewSessionRepository(db)
	sessionService := services.NewSessionService(tokenService, sessionRepository)

	authMiddleware := middlewares.NewAuthMiddleware(ecdsa, requestContext, sessionService)

	sessionHandler := handlers.NewSessionHandler(requestContext, sessionService)

	router.GET("/me/sessions", authMiddleware.Authenticated(sessionHandler.GetUserSessions))
	router.DELETE("/me/sessions/{sessionId}", authMiddleware.Authenticated(sessionHandler.RevokeSession))
	router.DELETE("/me/sessions", authMiddleware.Authenticated(sessionHandler.RevokeAllSessions))
}

func setupPostRoutes(db *sql.DB, router *Router) {
	ecdsa := pkgs.NewEcdsaKeyPair()
	requestContext := pkgs.NewRequestContext()

	tokenService := services.NewTokenService(ecdsa)
	sessionRepository := repositories.NewSessionRepository(db)
	sessionService := services.NewSessionService(tokenService, sessionRepository)

	authMiddleware := middlewares.NewAuthMiddleware(ecdsa, requestContext, sessionService)

	postRepository := repositories.NewPostRepository(db)
	likeRepository := repositories.NewLikeRepository(db)
	userRepository := repositories.NewUserRepository(db)
	likeService := services.NewLikeService(likeRepository)
	postService := services.NewPostService(likeService, postRepository, userRepository)
	postHandler := handlers.NewPostHandler(requestContext, postService)

	router.POST("/posts", authMiddleware.Authenticated(postHandler.CreatePost))
	router.GET("/posts/{postId}", authMiddleware.Authenticated(postHandler.GetPostByID))
	router.PUT("/posts/{postId}", authMiddleware.Authenticated(postHandler.UpdatePost))
	router.DELETE("/posts/{postId}", authMiddleware.Authenticated(postHandler.DeletePost))
	router.POST("/posts/{postId}/like", authMiddleware.Authenticated(postHandler.LikePost))
	router.POST("/posts/{postId}/unlike", authMiddleware.Authenticated(postHandler.UnlikePost))
	router.GET("/me/posts", authMiddleware.Authenticated(postHandler.GetPostsByAuthorID))
	router.GET("/users/{username}/posts", authMiddleware.Authenticated(postHandler.GetPostsByUsername))
}

func setupFeedRoutes(db *sql.DB, router *Router) {
	ecdsa := pkgs.NewEcdsaKeyPair()
	requestContext := pkgs.NewRequestContext()

	tokenService := services.NewTokenService(ecdsa)
	sessionRepository := repositories.NewSessionRepository(db)
	sessionService := services.NewSessionService(tokenService, sessionRepository)

	authMiddleware := middlewares.NewAuthMiddleware(ecdsa, requestContext, sessionService)

	likeRepository := repositories.NewLikeRepository(db)
	likeService := services.NewLikeService(likeRepository)

	feedRepository := repositories.NewFeedRepository(db)

	feedService := services.NewFeedService(likeService, feedRepository)

	feedHandler := handlers.NewFeedHandler(requestContext, feedService)

	router.GET("/feed", authMiddleware.Authenticated(feedHandler.GetFeed))
}
