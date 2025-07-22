package errors

import "errors"

var (
	ErrBookListEmpty  = errors.New("book list is empty")
	ErrBookIDNotFound = errors.New("book ID not found")
	ErrBookNotFound = errors.New("book not found")
)
