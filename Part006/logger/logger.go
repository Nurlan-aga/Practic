package logger

type LoggerAPI struct {
	infologs  []string
	warnlogs  []string
	fatallogs []string
}

func (l *LoggerAPI) Info(msg IStringer) {
	l.infologs = append(l.infologs, msg.String())
}
func (l *LoggerAPI) Warn(msg IStringer) {
	l.warnlogs = append(l.warnlogs, msg.String())
}
func (l *LoggerAPI) Fatal(msg IStringer) {
	l.fatallogs = append(l.fatallogs, msg.String())
}
func (l *LoggerAPI) ReadLogs() []string {
	allLogs := []string{}
	allLogs = append(allLogs, l.infologs...)
	allLogs = append(allLogs, l.warnlogs...)
	allLogs = append(allLogs, l.fatallogs...)
	return allLogs
}
