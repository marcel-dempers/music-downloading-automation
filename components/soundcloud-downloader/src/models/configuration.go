package models
type Route struct {
	Name string
}

type Configuration struct {
	Routes []Route
	Environments Environments
}

type Environments struct {
    Prod Environment
	Dev Environment
	ContainerRegistryUrl string
	ContainerRepository string
}

type Environment struct {
	Clusters []Cluster
}

type Cluster struct {
    Name string
}