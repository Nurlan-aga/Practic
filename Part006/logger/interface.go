package logger

type IStringer interface {
	String() string
}

type ILogger interface {
	Info(msg IStringer)
	Warn(msg IStringer)
	Fatal(msg IStringer)
	ReadLogs() []string
}
