package models

type Songlist struct {
	Total_Rows	int 		`json:"total_rows"`
	Offset	int 			`json:"offset"`
	Rows []Song				`json:"rows"`
}

type SonglistItem struct {
	ID  string 				`json:"id"`
	Key  string 			`json:"key"`
	Value Song 				`json:"value"`
}

type Song struct {
	ID  string 				`json:"id"`
	Uploader string			`json:"uploader"`
	Url string				`json:"webpage_url"`
	Track string			`json:"track"`
	Artist string 			`json:"artist"`
	Extractor string		`json:"extractor"`
	Title string 			`json:"title"`
	License string 			`json:"license"`
}