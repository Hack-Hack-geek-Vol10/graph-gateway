package config

var Config = &config{}

type config struct {
	Server  Server
	Service Service
}

type Server struct {
	Port string `env:"SERVER_ADDR" envDefault:":8080"`
}

type Service struct {
	UserServiceAddr string `env:"USER_SERVICE_ADDR" envDefault:"localhost:8081"`
}
