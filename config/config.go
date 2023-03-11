package config

type config interface {
	HOST_URL() string
	HOST_PORT() string
	MAX_QUEUE_WORKER() int
	REDIS_HOST() string
	REDIS_PORT() string
	REDIS_USERNAME() string
	REDIS_PASSWORD() string
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

func MAX_QUEUE_WORKER() int {
	return c.MAX_QUEUE_WORKER()
}

func REDIS_HOST() string {
	return c.REDIS_HOST()
}

func REDIS_PORT() string {
	return c.REDIS_PORT()
}

func REDIS_USERNAME() string {
	return c.REDIS_USERNAME()
}

func REDIS_PASSWORD() string {
	return c.REDIS_PASSWORD()
}
