package comment

import "github.com/uptrace/bun"

type Create struct {
	bun.BaseModel `bun:"table:comments"`
	ID           int    `json:"id" bun:"id,pk,autoincrement"`
	Text         string `json:"text" bun:"text"`
	TaskID       int    `json:"task_id" bun:"task_id"`
	ProgrammerID string `json:"programmer_id" bun:"programmer_id"`
	CreatedAt    string `json:"created_at" bun:"created_at,default:current_timestamp"`
}

type List struct {
	Count int
	L []Create
}