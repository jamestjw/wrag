package wrag

func Initialise(configPath string) {
	initialiseConfig(configPath)
	fetchToken()
}
