package entity

type User struct {
	ID        string `json:"id" bun:"id,pk"`
	FullName  string `json:"full_name"`
	Password  string `json:"password"`
	Avatar    string `json:"avatar"`
	Role      string `json:"role"`
	BirthDate string `json:"birth_date"`
	Phone     string `json:"phone"`
	Position  string `json:"position"`
}

type UserAttendance struct {
	Type      string `json:"type"`
	UserID    string `json:"user_id"`
	CreatedAt string `json:"created_at"`
}

type UserList struct {
	Count int
	L []User
}