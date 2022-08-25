package models

type Collection = string
type Size = int

const (
	BOOK_COLLECTION Collection = "books"
	MAX_SIZE        Size       = 2 * 1024 * 1024
)
