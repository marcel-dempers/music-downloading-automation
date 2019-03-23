package models
type Test struct {
	Name string
}

type Configuration struct {
	TestCollection []Test
	RabbitMq RabbitMq
}

type RabbitMq struct {
	Host string
	Port int
	Username string
	Password string
}