package config

var Config = &config{}

type config struct {
	Server  Server
	Service Service
}

type Server struct {
	Port string `env:"SERVER_ADDR" envDefault:"8080"`
}

type Service struct {
	UserServiceAddr    string `env:"USER_SERVICE_ADDR" envDefault:"localhost:8081"`
	ProjectServiceAddr string `env:"PROJECT_SERVICE_ADDR" envDefault:"localhost:8082"`
	MemberServiceAddr  string `env:"MEMBER_SERVICE_ADDR" envDefault:"localhost:8083"`
	ImageServiceAddr   string `env:"IMAGE_SERVICE_ADDR" envDefault:"localhost:8084"`
	//TokenServiceAddr   string `env:"TOKEN_SERVICE_ADDR" envDefault:"localhost:8085"`
}
