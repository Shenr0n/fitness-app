CREATE TABLE "default_exercises" (
  "de_id" bigserial PRIMARY KEY,
  "exercise_name" varchar,
  "muscle_group" varchar
);

CREATE TABLE "default_workouts" (
  "dw_id" bigserial PRIMARY KEY,
  "workout_name" varchar
);

CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "dob" varchar NOT NULL,
  "weight" integer NOT NULL,
  "height" integer NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "user_details" (
  "ud_id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "age" integer NOT NULL,
  "weight" integer NOT NULL,
  "height" integer NOT NULL,
  "goal_weight" integer NOT NULL,
  "diet_pref" varchar NOT NULL,
  "food_allergies" varchar NOT NULL,
  "daily_cal_intake_goal" integer NOT NULL,
  "activity_level" varchar NOT NULL,
  "current_fitness" varchar NOT NULL,
  "fitness_goal" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "user_track" (
  "ut_id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "weight" integer NOT NULL,
  "ut_date" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "user_macros" (
  "um_id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "calories" integer NOT NULL,
  "fats" integer NOT NULL,
  "protein" integer NOT NULL,
  "carbs" integer NOT NULL,
  "um_date" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "exercises" (
  "exer_id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "exercise_name" varchar NOT NULL,
  "muscle_group" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "workouts" (
  "workout_id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "workout_name" varchar UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "user_track_workouts" (
  "utw_id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "workout_id" bigint NOT NULL,
  "workout_name" varchar NOT NULL,
  "utw_date" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "workout_exercises" (
  "we_id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "workout_id" bigint NOT NULL,
  "exer_id" bigint NOT NULL,
  "weights" int NOT NULL,
  "sets" int NOT NULL,
  "reps" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "user_track" ("username", "ut_date");

CREATE INDEX ON "user_macros" ("username", "um_date");

CREATE INDEX ON "exercises" ("username");

CREATE UNIQUE INDEX ON "exercises" ("username", "exercise_name");

CREATE UNIQUE INDEX ON "workouts" ("username", "workout_name");

CREATE INDEX ON "user_track_workouts" ("username", "utw_date");

CREATE UNIQUE INDEX ON "workout_exercises" ("username", "workout_id", "exer_id");

COMMENT ON COLUMN "users"."weight" IS 'must be positive';

COMMENT ON COLUMN "users"."height" IS 'must be positive';

COMMENT ON COLUMN "user_track"."weight" IS 'must be positive';

COMMENT ON COLUMN "user_macros"."calories" IS 'must be positive';

COMMENT ON COLUMN "user_macros"."fats" IS 'must be positive';

COMMENT ON COLUMN "user_macros"."protein" IS 'must be positive';

COMMENT ON COLUMN "user_macros"."carbs" IS 'must be positive';

COMMENT ON COLUMN "workout_exercises"."weights" IS 'must be positive';

COMMENT ON COLUMN "workout_exercises"."sets" IS 'must be positive';

COMMENT ON COLUMN "workout_exercises"."reps" IS 'must be positive';

ALTER TABLE "user_details" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "user_track" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "user_macros" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "exercises" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "workouts" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "user_track_workouts" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "user_track_workouts" ADD FOREIGN KEY ("workout_id") REFERENCES "workouts" ("workout_id");

ALTER TABLE "user_track_workouts" ADD FOREIGN KEY ("workout_name") REFERENCES "workouts" ("workout_name");

ALTER TABLE "workout_exercises" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "workout_exercises" ADD FOREIGN KEY ("workout_id") REFERENCES "workouts" ("workout_id");

ALTER TABLE "workout_exercises" ADD FOREIGN KEY ("exer_id") REFERENCES "exercises" ("exer_id");


INSERT INTO default_exercises (exercise_name, muscle_group) VALUES
('Push-ups', 'Chest, Shoulders, Triceps'),
('Pull-ups', 'Back, Biceps'),
('Rows', 'Back, Biceps'),
('Running', 'Legs, Core');

INSERT INTO default_workouts (workout_name) VALUES
('Push'),
('Pull'),
('Legs'),
('Cardio');