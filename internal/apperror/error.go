package apperror

const (
	ErrNotFound      Error = "object not found"
	ErrAlreadyExists Error = "such user already exists"
)

type Error string

func (e Error) Error() string {
	return string(e)
}
