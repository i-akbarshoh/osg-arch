package user

type Filter struct {
	Limit     *int
	Offset    *int
	FirstName *string
}

type Create struct {
	FirstName string `json:"first_name" form:""`
}

type Update struct {
}

type List struct {
	Id        int
	FirstName string
}

type Detail struct {
}
