package calculator

type Simple struct {
	Logger
}

func NewSimple(logger Logger) Simple {
	return Simple{logger}
}
