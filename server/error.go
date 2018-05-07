package server

type sError struct {
	Code	uint
	Msg	string
}

func (this *sError)Error() string {
	return this.Msg
}
