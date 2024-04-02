package api

type getUserRequest struct {
	Username string `uri:"username" binding:"required,min=1"`
}

type getPageRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

type getUserAndDateRequest struct {
	Username string `uri:"username" binding:"required,min=1"`
	UmDate   string `uri:"um_date" binding:"required,min=1"`
}

type createUserRequest struct {
	Username       string `json:"username" binding:"required,alphanum"`
	HashedPassword string `json:"password" binding:"required,min=6"`
	FullName       string `json:"full_name" binding:"required"`
	Email          string `json:"email" binding:"required,email"`
	Dob            string `json:"dob" binding:"required" time_format:"2006-01-02"`
	Weight         int32  `json:"weight" binding:"required"`
	Height         int32  `json:"height" binding:"required"`
}

type userResponse struct {
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Dob      string `json:"dob"`
	Weight   int32  `json:"weight"`
	Height   int32  `json:"height"`
}

type recordMacrosRequest struct {
	Calories int32  `json:"calories" binding:"required"`
	Fats     int32  `json:"fats" binding:"required"`
	Protein  int32  `json:"protein" binding:"required"`
	Carbs    int32  `json:"carbs" binding:"required"`
	UmDate   string `json:"um_date" binding:"required" time_format:"2006-01-02"`
}

type recordUserTrackRequest struct {
	Weight int32  `json:"weight" binding:"required"`
	UtDate string `json:"ut_date" binding:"required"`
}

type trackResponse struct {
	Weight int32  `json:"weight"`
	UtDate string `json:"ut_date"`
}

type createExerciseRequest struct {
	ExerciseName string `json:"exercise_name" binding:"required"`
	MuscleGroup  string `json:"muscle_group" binding:"required"`
}

type deleteExerciseRequest struct {
	ExerID int64 `json:"exer_id" binding:"required"`
}

type deleteWorkoutRequest struct {
	WorkoutID int64 `json:"workout_id" binding:"required"`
}

type createWorkoutRequest struct {
	WorkoutName string `json:"workout_name" binding:"required"`
}
