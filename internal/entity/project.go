package entity

type Project struct {
	ID           int `json:"id" bun:"id,autoincrement,pk"`
	Name         string `json:"name"`
	StartedDate  string `json:"started_date" bun:"default:current_timestamp"`
	FinishedDate string `json:"finished_date" bun:"finished_date,nullzero"`
	Status       string `json:"status"`
	TeamLeadID   string `json:"teamlead_id" bun:"teamlead_id"`
	Attachment   string `json:"attachment"`
}

type List struct {
	Count int `json:"count" bun:"count"`
	Projects []Project
}
