package constants

import "os"

type appEnvironment struct {
	/* The default value is 80 or override from environment */
	Port    string
	/* The default value is 192.168.1.255 or override from environment */
	Network string
}

var AppEnv appEnvironment

func Load() {
	_port := os.Getenv("PORT")
	if _port == "" {
		_port = "80"
	}

	_network := os.Getenv("NETWORK")
	if _network == "" {
		_network = "192.168.1.255"
	}

	AppEnv = appEnvironment{
		Port: _port,
		Network: _network,
	}
}
