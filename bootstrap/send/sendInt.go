package send
type SendInter interface {
	Send(message SendMessage) error
}

type SendMessage struct {
	SendAddress string
	Message string
}
