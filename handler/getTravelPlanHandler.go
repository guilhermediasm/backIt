package handler

import (
	"back-It/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Get travel
// @Description Get a travel in description
// @Tags Openings
// @Accept json
// @Produce json
// @Param request body CreateOpeningRequest true "Request body"
// @Success 200 {object} CreateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [post]

func GetTravelPlan(ctx *gin.Context) {
	// Lógica para obter um plano de viagem com base nas solicitações do usuário
	// e interações com chatGPT

	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	travel := models.TravelPlan{
		ID:          1,
		Destination: "Paris, França",
		StartDate:   "2023-10-15",
		EndDate:     "2023-10-22",
		// Preencha outros campos do plano de viagem conforme necessário
	}

	// Converte o plano de viagem em JSON
	responseJSON, err := json.Marshal(travel)
	if err != nil {
		http.Error(w, "Erro ao serializar os dados", http.StatusInternalServerError)
		return
	}

	// Define o cabeçalho Content-Type para JSON
	w.Header().Set("Content-Type", "application/json")

	// Escreve a resposta JSON
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}
