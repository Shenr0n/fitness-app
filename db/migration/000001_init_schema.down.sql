ALTER TABLE "workout_exercises" DROP CONSTRAINT IF EXISTS "workout_exercises_username_fkey";
ALTER TABLE "workout_exercises" DROP CONSTRAINT IF EXISTS "workout_exercises_workout_id_fkey";
ALTER TABLE "workout_exercises" DROP CONSTRAINT IF EXISTS "workout_exercises_exer_id_fkey";

ALTER TABLE "user_track_workouts" DROP CONSTRAINT IF EXISTS "user_track_workouts_username_fkey";
ALTER TABLE "user_track_workouts" DROP CONSTRAINT IF EXISTS "user_track_workouts_workout_id_fkey";
ALTER TABLE "user_track_workouts" DROP CONSTRAINT IF EXISTS "user_track_workouts_workout_name_fkey";

ALTER TABLE "workouts" DROP CONSTRAINT IF EXISTS "workouts_username_fkey";

ALTER TABLE "exercises" DROP CONSTRAINT IF EXISTS "exercises_username_fkey";

ALTER TABLE "user_macros" DROP CONSTRAINT IF EXISTS "user_macros_username_fkey";

ALTER TABLE "user_track" DROP CONSTRAINT IF EXISTS "user_track_username_fkey";


DROP TABLE IF EXISTS "workout_exercises";
DROP TABLE IF EXISTS "user_track_workouts";
DROP TABLE IF EXISTS "workouts";
DROP TABLE IF EXISTS "exercises";
DROP TABLE IF EXISTS "user_macros";
DROP TABLE IF EXISTS "user_track";
DROP TABLE IF EXISTS "users";

DROP TABLE IF EXISTS "default_exercises";
DROP TABLE IF EXISTS "default_workouts";