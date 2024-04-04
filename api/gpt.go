package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/Shenr0n/fitness-app/token"
	"github.com/gin-gonic/gin"
	openai "github.com/sashabaranov/go-openai"
)

func (server *Server) gptChat(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if req.Username != authPayload.Username {
		ctx.JSON(http.StatusUnauthorized, nil)
		return
	}

	var reqBody getGPTQuestionRequest
	var gptReqBody getGPTRequest

	if err := ctx.BindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var tempString string
	if reqBody.QuestionType == "diet" {
		dietInfo, err := server.store.GetDietDetails(ctx, authPayload.Username)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		tempString =
			fmt.Sprintf("Name: %s, Age: %d, Weight: %d kg, Height: %d cm, Goal weight: %d kg, Diet preference: %s, Food allegies: %s, Daily Calories Goal: %d",
				authPayload.Username, dietInfo.Age, dietInfo.Weight, dietInfo.Height, dietInfo.GoalWeight, dietInfo.DietPref, dietInfo.FoodAllergies, dietInfo.DailyCalIntakeGoal)
	} else if reqBody.QuestionType == "fitness" {
		fitnessInfo, err := server.store.GetFitnessDetails(ctx, authPayload.Username)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		tempString =
			fmt.Sprintf("Name: %s, Age: %d, Weight: %d kg, Height: %d cm, Goal weight: %d kg, Current Fitness: %s, Activity Level: %s, Fitness goal: %s",
				authPayload.Username, fitnessInfo.Age, fitnessInfo.Weight, fitnessInfo.Height, fitnessInfo.GoalWeight, fitnessInfo.CurrentFitness, fitnessInfo.ActivityLevel, fitnessInfo.FitnessGoal)

	}
	gptReqBody.ReqText = fmt.Sprintf("%s Based on this, answer the following in highly condensed format. %s", tempString, reqBody.ReqText)

	// OpenAI token
	client := openai.NewClient(os.Getenv("OPENAI_KEY"))

	ctxAI, cancel := context.WithCancel(ctx)
	defer cancel()

	requestData := openai.CompletionRequest{
		Model:     openai.GPT3Dot5TurboInstruct,
		MaxTokens: 300,
		Prompt:    string(gptReqBody.ReqText),
	}
	fmt.Println("Request sent")

	resp, err := client.CreateCompletion(ctxAI, requestData)

	//fmt.Println(resp.Choices[0].Text)

	if err != nil {
		if openaiError, ok := err.(*openai.APIError); ok {
			fmt.Println("OpenAI Error:")
			ctx.JSON(http.StatusInternalServerError, openaiError.Error())
			return
		} else {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
		}
	}
	cleanedResp := strings.ReplaceAll((resp.Choices[0].Text), "\n", "")
	jsonResponse := getGPTResponse{
		RespText: cleanedResp,
	}
	ctx.JSON(http.StatusOK, jsonResponse)
}
