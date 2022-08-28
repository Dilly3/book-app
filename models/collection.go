package models

type Collection = string
type Size = int
type Role = string

const (
	BOOK_COLLECTION Collection = "books"
	USER_COLLECTION Collection = "users"
	MAX_SIZE        Size       = 2 * 1024 * 1024
	DEFAULT_COST    int        = 10
)

const (
	LIBRARY_USER Role = "library_user"
	ADMIN_USER   Role = "admin"
)
