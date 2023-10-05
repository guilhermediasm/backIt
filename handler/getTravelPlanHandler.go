package handler

import (
	"context"
	"encoding/json"
	"log"

	"github.com/ayush6624/go-chatgpt"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Get travel
// @Description Get a travel in description
// @Tags Openings
// @Accept json
// @Produce json
// @Param request body CreateTravelPlanRequest true "Request body"
// @Success 200 {object} CreateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [post]

func GetTravelPlan(ctx *gin.Context) {
	// Lógica para obter um plano de viagem com base nas solicitações do usuário
	// e interações com chatGPT

	// key := os.Getenv("OPENAI_KEY")
	key := "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"
	c, err := chatgpt.NewClient(key)
	if err != nil {
		log.Fatal(err)
	}

	ctxChat := context.Background()
	res, err := c.SimpleSend(ctxChat, "Hey, Explain GoLang to me in 2 sentences.")
	if err != nil {
		log.Fatal(err)
	}

	a, _ := json.MarshalIndent(res, "", "  ")
	log.Println(string(a))

	res, err = c.Send(ctxChat, &chatgpt.ChatCompletionRequest{
		Model: chatgpt.GPT4,
		Messages: []chatgpt.ChatMessage{
			{
				Role:    chatgpt.ChatGPTModelRoleSystem,
				Content: "Hey, Explain GoLang to me in 2 sentences.",
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	a, _ = json.MarshalIndent(res, "", "  ")
	log.Println(string(a))

	sendSuccess(ctx, "travelplan-created", string(a))
}
