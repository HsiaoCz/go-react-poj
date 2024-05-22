package storage

type Store struct {
	User UserStorer
	Todo TodoStorer
}
