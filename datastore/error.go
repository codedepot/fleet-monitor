package datastore

type StoreError interface {
	Msg() string
}

type NotFoundError struct {
	msg string
}

type ServerError struct {
	msg string
}

func (e *NotFoundError) Msg() string {
	return e.msg
}

func (e *ServerError) Msg() string {
	return e.msg
}
