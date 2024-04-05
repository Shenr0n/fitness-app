// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type DefaultExercise struct {
	DeID         int64          `json:"de_id"`
	ExerciseName sql.NullString `json:"exercise_name"`
	MuscleGroup  sql.NullString `json:"muscle_group"`
}

type DefaultWorkout struct {
	DwID        int64          `json:"dw_id"`
	WorkoutName sql.NullString `json:"workout_name"`
}

type Exercise struct {
	ExerID       int64     `json:"exer_id"`
	Username     string    `json:"username"`
	ExerciseName string    `json:"exercise_name"`
	MuscleGroup  string    `json:"muscle_group"`
	CreatedAt    time.Time `json:"created_at"`
}

type Session struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type User struct {
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
	Dob            string `json:"dob"`
	// must be positive
	Weight int32 `json:"weight"`
	// must be positive
	Height    int32     `json:"height"`
	CreatedAt time.Time `json:"created_at"`
}

type UserDetail struct {
	UdID               int64     `json:"ud_id"`
	Username           string    `json:"username"`
	Age                int32     `json:"age"`
	Weight             int32     `json:"weight"`
	Height             int32     `json:"height"`
	GoalWeight         int32     `json:"goal_weight"`
	DietPref           string    `json:"diet_pref"`
	FoodAllergies      string    `json:"food_allergies"`
	DailyCalIntakeGoal int32     `json:"daily_cal_intake_goal"`
	ActivityLevel      string    `json:"activity_level"`
	CurrentFitness     string    `json:"current_fitness"`
	FitnessGoal        string    `json:"fitness_goal"`
	CreatedAt          time.Time `json:"created_at"`
}

type UserMacro struct {
	UmID     int64  `json:"um_id"`
	Username string `json:"username"`
	// must be positive
	Calories int32 `json:"calories"`
	// must be positive
	Fats int32 `json:"fats"`
	// must be positive
	Protein int32 `json:"protein"`
	// must be positive
	Carbs     int32     `json:"carbs"`
	UmDate    string    `json:"um_date"`
	CreatedAt time.Time `json:"created_at"`
}

type UserTrack struct {
	UtID     int64  `json:"ut_id"`
	Username string `json:"username"`
	// must be positive
	Weight    int32     `json:"weight"`
	UtDate    string    `json:"ut_date"`
	CreatedAt time.Time `json:"created_at"`
}

type UserTrackWorkout struct {
	UtwID       int64     `json:"utw_id"`
	Username    string    `json:"username"`
	WorkoutID   int64     `json:"workout_id"`
	WorkoutName string    `json:"workout_name"`
	UtwDate     string    `json:"utw_date"`
	CreatedAt   time.Time `json:"created_at"`
}

type Workout struct {
	WorkoutID   int64     `json:"workout_id"`
	Username    string    `json:"username"`
	WorkoutName string    `json:"workout_name"`
	CreatedAt   time.Time `json:"created_at"`
}

type WorkoutExercise struct {
	WeID      int64  `json:"we_id"`
	Username  string `json:"username"`
	WorkoutID int64  `json:"workout_id"`
	ExerID    int64  `json:"exer_id"`
	// must be positive
	Weights int32 `json:"weights"`
	// must be positive
	Sets int32 `json:"sets"`
	// must be positive
	Reps      int32     `json:"reps"`
	CreatedAt time.Time `json:"created_at"`
}
