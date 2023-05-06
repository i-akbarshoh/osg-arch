package user

type Filter struct {
	Limit     *int
	Offset    *int
	FirstName *string
}

type Create struct {
	ID        string `json:"id" bun:"id"`
	FullName  string `json:"full_name" bun:"full_name"`
	Password  string `json:"password" bun:"password"`
	Avatar    string `json:"avatar" bun:"avatar"`
	Role      string `json:"role" bun:"role"`
	BirthDate string `json:"birth_date" bun:"birth_date"`
	Phone     string `json:"phone" bun:"phone"`
	Position  string `json:"position" bun:"position"`
}

type Get struct {
	ID        string `json:"id" bun:"id"`
	FullName  string `json:"full_name" bun:"full_name"`
	Avatar    string `json:"avatar" bun:"avatar"`
	Role      string `json:"role" bun:"role"`
	BirthDate string `json:"birth_date" bun:"birth_date"`
	Phone     string `json:"phone" bun:"phone"`
	Position  string `json:"position" bun:"position"`
}

type List []Get

type Delete struct {
	Phone    string `json:"phone" bun:"phone"`
	Password string `json:"password" bun:"password"`
}

type Update struct {
	ID        string `json:"id" bun:"id"`
	FullName  string `json:"full_name" bun:"full_name"`
	Avatar    string `json:"avatar" bun:"avatar"`
	Role      string `json:"role" bun:"role"`
	BirthDate string `json:"birth_date" bun:"birth_date"`
	Phone     string `json:"phone" bun:"phone"`
	Position  string `json:"position" bun:"position"`
}
