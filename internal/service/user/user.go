package user

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}
