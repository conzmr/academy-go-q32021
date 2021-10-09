package controller

type AppController struct {
	Coach    interface{ CoachController }
	Exercise interface{ ExerciseController }
}
