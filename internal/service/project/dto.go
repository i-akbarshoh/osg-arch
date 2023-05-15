package project

type Create struct {
	Name         string `json:"name" bun:"name"`
	Status       string `json:"status" bun:"status"`
	TeamLeadID   string `json:"teamlead_id" bun:"teamlead_id"`
	Attachment   string `json:"attachment" bun:"attachment"`
}

type List struct {
	ID           string `json:"id" bun:"id,autoincrement"`
	Name         string `json:"name"`
	StartedDate  string `json:"started_date" bun:"default:current_timestamp"`
	FinishedDate string `json:"finished_date" bun:"finished_date,nullzero"`
	Status       string `json:"status"`
	TeamLeadID   string `json:"teamlead_id" bun:"teamlead_id"`
	Attachment   string `json:"attachment"`
}

type Update struct {
	ID int `bun:"id,pk"`
	Status string `bun:"status"`
}