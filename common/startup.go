package common

// StartUp bootstrapps the application
func StartUp() {
	// Initialize AppConfig variable
	initConfig()
	// Initialize Logger objects with Log Level
	setLogLevel(Level(AppConfig.LogLevel))
}
