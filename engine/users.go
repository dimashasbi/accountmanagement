package engine

type (
	// Users is the interface for interactor
	Users interface {
	}

	usersStr struct {
		repository UsersRepository
	}
)

func (f *engineFactory) NewUsersEngine() usersStr {
	return &usersStr{
		repository: f.NewUsersRespository(),
	}
}
