package constants

import (
	"os"
	"strconv"

	"github.com/google/uuid"
)

type appEnvironment struct {
	/* The default value is 80 or override from environment */
	Port string
	/* The default value is 192.168.1.255 or override from environment */
	Network string
	/* The default value is 3600 * 24 (24h) or override from environment */
	CookieTTL    int
	CookieSecret string
	Password     string
}

var AppEnv appEnvironment

func (c *appEnvironment) Load() {
	_port := os.Getenv("PORT")
	if _port == "" {
		_port = "80"
	}

	_network := os.Getenv("NETWORK")
	if _network == "" {
		_network = "192.168.1.255"
	}

	_defaultCookieTTL := 3600 * 24
	_cookieTTL := os.Getenv("COOKIE_TTL")
	if _cookieTTL != "" {
		parsedVal, err := strconv.Atoi(_cookieTTL)
		if err == nil {
			_defaultCookieTTL = parsedVal
		}
	}

	_cookieSecret := os.Getenv("COOKIE_SECRET")
	if _cookieSecret == "" {
		_cookieSecret = uuid.New().String()
	}

	c.Port = _port
	c.Network = _network
	c.CookieTTL = _defaultCookieTTL
	c.CookieSecret = _cookieSecret
	c.Password = os.Getenv("PASSWORD")
}
