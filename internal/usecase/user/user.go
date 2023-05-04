package user

type UseCase struct {
	u User
}

func NewUseCase(u User) *UseCase {
	return &UseCase{
		u: u,
	}
}
