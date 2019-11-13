package engine

type (
	// SystemSettingsRepository defines the methods that any
	// data storage provider needs to implement to get
	// and store greetings
	UsersRepository interface {
	}

	// StorageFactory is the interface that a storage
	// provider needs to implement so that the engine can
	// request repository instances as it needs them
	StorageFactory interface {
		// NewParameterRepository returns a storage specific
		NewUsersRespository() UsersRepository
	}
)
