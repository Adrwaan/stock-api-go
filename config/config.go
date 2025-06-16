package config

func Init() {

}

func GetLogger() *Logger {
	logger := NewLogger("")

	return logger
}
