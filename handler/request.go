package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CreateTravelPlanRequest struct {
	Location    string `json:"location"`
	Destination string `json:"destination"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
}

func (r *CreateTravelPlanRequest) Validate() error {
	if r.Destination == "" && r.StartDate == "" && r.Location == "" && r.EndDate == "" {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.Destination == "" {
		return errParamIsRequired("destination", "string")
	}
	if r.StartDate == "" {
		return errParamIsRequired("startDate", "string")
	}
	if r.Location == "" {
		return errParamIsRequired("location", "string")
	}
	if r.EndDate == "" {
		return errParamIsRequired("endDate", "string")
	}
	return nil
}
