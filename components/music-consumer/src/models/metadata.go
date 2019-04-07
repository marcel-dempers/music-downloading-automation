package models

type MetadataFile struct {
	ID  string 				`json:"id"`
	Uploader string			`json:"uploader"`
	Url string				`json:"webpage_url"`
	Track string			`json:"track"`
	Artist string 			`json:"artist"`
	Extractor string		`json:"extractor"`
	Title string 			`json:"title"`
	License string 			`json:"license"`
}