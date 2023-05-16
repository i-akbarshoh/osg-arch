package task

import "github.com/uptrace/bun"

type Create struct {
	bun.BaseModel `bun:"table:tasks"`
	Title         string `bun:"title"`
	Description   string `bun:"description"`
	StartAt       string `bun:"start_at"`
	FinishAt      string `bun:"finish_at"`
	ProjectID     int    `bun:"project_id"`
	UserID        string `bun:"programmer_id"`
}

type Update struct {
	bun.BaseModel `bun:"table:tasks"`
	ID            int    `bun:"id,pk"`
	Status        string `bun:"status"`
}

type Get struct {
	bun.BaseModel `bun:"table:tasks"`
	ID            int    `json:"id" bun:"id,pk"`
	Title         string `bun:"title"`
	Description   string `bun:"description"`
	StartAt       string `bun:"start_at"`
	FinishAt      string `bun:"finish_at"`
	ProjectID     int    `bun:"project_id"`
	UserID        string `bun:"programmer_id"`
	StartedAt     string `json:"started_at" bun:"started_at"`
	FinishedAt    string `json:"finished_at" bun:"finished_at"`
	Status        string `json:"status" bun:"status"`
	Attachment    string `json:"attachment" bun:"attachment"`
}

type List struct {
	Count int
	List []Get
}