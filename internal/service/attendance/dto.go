package attendance

import "github.com/uptrace/bun"

type Create struct {
	bun.BaseModel `bun:"table:attendance"`
	Type          string `json:"type" bun:"type"`
	UserID        string `json:"user_id" bun:"user_id"`
	CreatedAt     string `json:"created_at" bun:"date,default:current_timestamp"`
}

type List struct {
	Type      string `bun:"type"`
	CreatedAt string `json:"created_at" bun:"date,default:current_timestamp"`
}
