package models

type Collection = string
type Size = int
type Role = string
type Status = string

const (
	BOOK_COLLECTION Collection = "books"
	USER_COLLECTION Collection = "users"
	MAX_SIZE        Size       = 2 * 1024 * 1024
	DEFAULT_COST    int        = 10
	RENTED          Status     = "rented"
	AVAILABLE       Status     = "available"
	LIBRARY_USER    Role       = "library_user"
	ADMIN_USER      Role       = "admin"
)
