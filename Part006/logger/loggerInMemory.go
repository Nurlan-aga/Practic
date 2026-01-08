package logger

type LoggerInMemory struct {
	infoLogs  []string
	warnLogs  []string
	fatalLogs []string
}

func (l *LoggerInMemory) Info(msg IStringer) {
	l.infoLogs = append(l.infoLogs, msg.String())
}

func (l *LoggerInMemory) Warn(msg IStringer) {
	l.warnLogs = append(l.warnLogs, msg.String())
}

func (l *LoggerInMemory) Fatal(msg IStringer) {
	l.fatalLogs = append(l.fatalLogs, msg.String())
}

func (l *LoggerInMemory) ReadLogs() []string {
	var logs []string
	logs = append(logs, l.infoLogs...)
	logs = append(logs, l.warnLogs...)
	logs = append(logs, l.fatalLogs...)
	return logs
}
