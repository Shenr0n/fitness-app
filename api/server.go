package api

import (
	db "github.com/Shenr0n/fitness-app/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// Home
	router.POST("/signup", server.createUser)

	// User
	router.GET("/users/:username", server.getUser)
	router.DELETE("/users/:username/delete", server.deleteUser)
	router.PATCH("/users/:username/password", nil)
	router.PATCH("/users/:username/weight", nil)
	router.PATCH("/users/:username/height", nil)

	// User Macros
	router.POST("/users/:username/macros", server.recordMacros)
	router.GET("/users/:username/macros", server.getMacros)
	router.GET("/users/:username/macros/:um_date", server.getMacroByDate)

	// User Track
	router.POST("/users/:username/track", server.recordUserTrack)
	router.GET("/users/:username/track", server.getUserTrack)

	// Exercises
	router.POST("/users/:username/exercises", server.createExercise)
	router.GET("/users/:username/exercises", server.getExercises)
	router.DELETE("/users/:username/exercises/delete", server.deleteExercise)

	// Workouts
	router.POST("/users/:username/workouts", server.createWorkout)
	router.GET("/users/:username/workouts", server.getWorkouts)
	router.DELETE("/users/:username/workouts/delete", server.deleteWorkout)

	// Workout Exercises
	router.POST("/users/:username/workout/exercises", nil)
	router.GET("/users/:username/workout/exercises", nil)
	router.DELETE("/users/:username/workout/exercises/delete", nil)
	router.DELETE("/users/:username/workout/delete", nil)

	// User Workouts Record
	router.POST("/users/:username/records", nil)
	router.GET("/users/:username/records", nil)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
