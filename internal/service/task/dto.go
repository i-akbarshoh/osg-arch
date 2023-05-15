package task

type Create struct {
	ID string `bun:"id,pk"`
	Title string `bun:"title"`
	Description string `bun:"description"`
	StartAt string `bun:"start_at"`
	FinishAt string `bun:"finish_at"`
	ProjectID int `bun:"project_id"`
	UserID string `bun:"user_id"`
}