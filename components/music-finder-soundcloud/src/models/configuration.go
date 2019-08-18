package models
type Test struct {
	Name string
}

type Configuration struct {
	TestCollection []Test
	Soundcloud Soundcloud
}

type Soundcloud struct {
	ApiUrl string
    ClientID string
}