package api

import (
	"fmt"

	db "github.com/Shenr0n/fitness-app/db/sqlc"
	"github.com/Shenr0n/fitness-app/token"
	"github.com/Shenr0n/fitness-app/util"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config     util.Config
	store      *db.Store
	router     *gin.Engine
	tokenMaker token.Maker
}

func NewServer(config util.Config, store *db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("NewPasetoMaker failed: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker}

	server.setupRouter()
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// Add routes to router
func (server *Server) setupRouter() {
	router := gin.Default()

	// Home
	router.POST("/signup", server.createUser)
	router.POST("/login", server.loginUser)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	// User
	authRoutes.GET("/users/:username", server.getUser)
	authRoutes.DELETE("/users/:username/delete", server.deleteUser)
	router.PATCH("/users/:username/password", nil)
	router.PATCH("/users/:username/weight", nil)
	router.PATCH("/users/:username/height", nil)

	// User Macros
	authRoutes.POST("/users/:username/macros", server.recordMacros)
	authRoutes.GET("/users/:username/macros", server.getMacros)
	authRoutes.GET("/users/:username/macros/:um_date", server.getMacroByDate)

	// User Track
	authRoutes.POST("/users/:username/track", server.recordUserTrack)
	authRoutes.GET("/users/:username/track", server.getUserTrack)

	// Exercises
	authRoutes.POST("/users/:username/exercises", server.createExercise)
	authRoutes.GET("/users/:username/exercises", server.getExercises)
	authRoutes.DELETE("/users/:username/exercises/delete", server.deleteExercise)

	// Workouts
	authRoutes.POST("/users/:username/workouts", server.createWorkout)
	authRoutes.GET("/users/:username/workouts", server.getWorkouts)
	authRoutes.DELETE("/users/:username/workouts/delete", server.deleteWorkout)

	// Workout Exercises
	authRoutes.POST("/users/:username/workout/exercises", server.addExerciseToWorkout)
	authRoutes.GET("/users/:username/workout/exercises", nil)

	// User Workouts Record
	router.POST("/users/:username/records", nil)
	router.GET("/users/:username/records", nil)

	server.router = router

}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
