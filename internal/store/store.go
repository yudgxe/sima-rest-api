package store

type Store interface {
	User() UserRepository
	Auth() AuthRepository
}
