-- name: RecordDetails :one
INSERT into user_details (
    username,
    age,
    weight,
    height,
    goal_weight,
    diet_pref,
    food_allergies,
    daily_cal_intake_goal,
    activity_level,
    current_fitness,
    fitness_goal
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
) 
RETURNING *;

-- name: GetDietDetails :one
SELECT age, weight, height, goal_weight, diet_pref, food_allergies, daily_cal_intake_goal FROM user_details
WHERE username = $1;

-- name: GetFitnessDetails :one
SELECT age, weight, height, goal_weight, activity_level, current_fitness, fitness_goal FROM user_details
WHERE username = $1;