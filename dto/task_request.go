package dto

type TaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status" enum:"TODO,DOING,DONE"`
}
