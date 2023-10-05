package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Create opening
// @Description Create a new jon opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param request body CreateTravelPlanRequest true "Request body"
// @Success 200 {object} CreateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [post]

func CreateTravelPlan(ctx *gin.Context) {
	// Lógica para criar um novo plano de viagem com base nas preferências do usuário
	// e interações com chatGPT
	request := CreateTravelPlanRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
}
