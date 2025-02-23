package sql

type PageResult[T any] struct {
	Total int
	List  T
}
