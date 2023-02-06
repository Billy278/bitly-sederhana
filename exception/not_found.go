package exception

type NotFound struct {
	Error string
}

func NewNotFound(error string) NotFound {
	return NotFound{
		Error: error,
	}
}
