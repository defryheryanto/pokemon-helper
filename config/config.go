package config

type config interface {
	HOST_URL() string
	HOST_PORT() string
}

var c config

func SetConfig(cfg config) {
	c = cfg
}

func HOST_URL() string {
	return c.HOST_URL()
}

func HOST_PORT() string {
	return c.HOST_PORT()
}
