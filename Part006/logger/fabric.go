package logger

const DRIVER_MEMORY = "memory"
const DRIVER_FILE = "file"

func NewLogger(driver string) ILogger {
	switch driver {
	case DRIVER_MEMORY:
		return &LoggerInMemory{}
	case DRIVER_FILE:
		return &LoggerFile{}
	default:
		return nil
	}
}
