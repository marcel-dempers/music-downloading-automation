package models
type Download struct {
	OutputPath string
}

type Configuration struct {
	Download Download
	RabbitMq RabbitMq
}

type RabbitMq struct {
	Host string
	Port int
	Username string
	Password string
}