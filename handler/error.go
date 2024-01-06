package handler

type Err struct {
	Msg string
}

func (e Err) Error() string {
	return e.Msg
}
