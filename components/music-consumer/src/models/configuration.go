package models
type Download struct {
	OutputPath string
}

type Configuration struct {
	Download Download
	RabbitMq RabbitMq
	Metadata Metadata
}

type RabbitMq struct {
	Host string
	Port int
	Username string
	Password string
}

type Metadata struct {
	License License
}

type License struct {
	NcsAutodetect bool
}