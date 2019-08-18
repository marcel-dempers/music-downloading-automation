package soundcloud

import (
	"net/http"
	"time"
	"app/models"
	"encoding/json"
	"io/ioutil"
	"strings"
)

type Track struct {
	Id int
	PermanentUrl string `json:"permalink_url"`
	License string
}

type Client struct {
	HttpClient http.Client
	ApiUrl string
	ClientID string
}

func(sc *Client) Init(config models.Configuration) *Client {
		
	sc.HttpClient = http.Client{
		Timeout: 10 * time.Second, //10 sec
	}
	
	sc.ClientID = config.Soundcloud.ClientID
	sc.ApiUrl = config.Soundcloud.ApiUrl

	return sc
}

func (sc *Client) GetRelatedTracks(id string) (trackList []Track, err error){

	req, err := http.NewRequest("GET", sc.ApiUrl + "/tracks/" + id + "/related?" + "client_id=" + sc.ClientID, nil)
	resp, err := sc.HttpClient.Do(req)

	if err != nil {
		return trackList, err
	}
	if resp != nil {
		defer resp.Body.Close()
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return trackList, err
		}

		err = json.Unmarshal(bodyBytes, &trackList)
		
		if err != nil {
			return trackList, err
		}

		filter := func(s Track) bool { 
			return strings.Contains(s.License,"cc-")
		}
		list := filterCCLicense(trackList, filter)
		
		return list, nil
	}
	return trackList, nil
    
}

func filterCCLicense(ss []Track, test func(Track) bool) (ret []Track) {
    for _, s := range ss {
        if test(s) {
            ret = append(ret, s)
        }
    }
    return
}
