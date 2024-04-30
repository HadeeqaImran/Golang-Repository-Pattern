package dto

type StatusRequest struct {
	Status string `json:"status" enum:"TODO,DOING,DONE"`
}
