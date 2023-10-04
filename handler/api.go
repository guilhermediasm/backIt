package handler

import (
	"back-It/models"
	"encoding/json"
	"net/http"
)

func GetTravelPlan(w http.ResponseWriter, r *http.Request) {
	// Lógica para obter um plano de viagem com base nas solicitações do usuário
	// e interações com chatGPT
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


