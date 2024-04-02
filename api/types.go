package api

type SuccessResponse struct {
	Msg string `json:"message"`
}
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
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Dob      string `json:"dob" binding:"required" time_format:"2006-01-02"`
	Weight   int32  `json:"weight" binding:"required,gt=0"`
	Height   int32  `json:"height" binding:"required,gt=0"`
}

type userResponse struct {
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Dob      string `json:"dob"`
	Weight   int32  `json:"weight"`
	Height   int32  `json:"height"`
}

type loginUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type loginUserResponse struct {
	AccessToken string       `json:"access_token"`
	User        userResponse `json:"user"`
}

type recordMacrosRequest struct {
	Calories int32  `json:"calories" binding:"required,gt=0"`
	Fats     int32  `json:"fats" binding:"required,gt=0"`
	Protein  int32  `json:"protein" binding:"required,gt=0"`
	Carbs    int32  `json:"carbs" binding:"required,gt=0"`
	UmDate   string `json:"um_date" binding:"required" time_format:"2006-01-02"`
}

type recordUserTrackRequest struct {
	Weight int32  `json:"weight" binding:"required,gt=0"`
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

type getExerciseRequest struct {
	ExerID int64 `json:"exer_id" binding:"required"`
}

type getWorkoutRequest struct {
	WorkoutID int64 `json:"workout_id" binding:"required"`
}

type getUtwIDRequest struct {
	UtwID int64 `json:"utw_id" binding:"required"`
}

type createWorkoutRequest struct {
	WorkoutName string `json:"workout_name" binding:"required"`
}

type addExerciseToWorkoutRequest struct {
	WorkoutID int64 `json:"workout_id" binding:"required"`
	ExerID    int64 `json:"exer_id" binding:"required"`
	Weights   int32 `json:"weights" binding:"required"`
	Sets      int32 `json:"sets" binding:"required"`
	Reps      int32 `json:"reps" binding:"required"`
}
type deleteExerciseInWorkoutRequest struct {
	WorkoutID int64 `json:"workout_id" binding:"required"`
	ExerID    int64 `json:"exer_id" binding:"required"`
}

type exerciseInWorkoutResponse struct {
	WorkoutID    int64  `json:"workout_id"`
	WorkoutName  string `json:"workout_name"`
	ExerID       int64  `json:"exer_id"`
	ExerciseName string `json:"exercise_name"`
	MuscleGroup  string `json:"muscle_group"`
	Weights      int32  `json:"weights"`
	Sets         int32  `json:"sets"`
	Reps         int32  `json:"reps"`
}

type recordUserWorkoutRequest struct {
	WorkoutID int64  `json:"workout_id" binding:"required"`
	UtwDate   string `json:"utw_date" binding:"required"`
}

type recordResponse struct {
	UtwID       int64  `json:"utw_id"`
	WorkoutID   int64  `json:"workout_id"`
	WorkoutName string `json:"workout_name"`
	UtwDate     string `json:"utw_date"`
}

type passwordChangeRequest struct {
	Password string `json:"password" binding:"required,min=6"`
}
