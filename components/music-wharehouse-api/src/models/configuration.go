package models

type Configuration struct {
	CouchDB CouchDB
}

type CouchDB struct {
	Host string
	Port int
	Database string
}