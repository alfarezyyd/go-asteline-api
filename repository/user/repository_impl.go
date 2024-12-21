package user

type RepositoryImpl struct{}

func NewUserRepositoryImpl() *RepositoryImpl {
	return &RepositoryImpl{}
}
