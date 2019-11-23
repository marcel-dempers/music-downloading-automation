package soundcloud

import (
	"net/http"
	"time"
	"app/models"
	"encoding/json"
	"io/ioutil"
	"strings"
	"strconv"
	"fmt"
	"net/url"
)

type Track struct {
	Id int
	PermanentUrl string `json:"permalink_url"`
	License string
}

type Resolve struct {
	Location string 
	Status string 
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

func (sc *Client)  GetTrackIdFromUrl(httpUrl string) (trackid int, err error) {
	
	fmt.Printf("httpUrl: %v\n", httpUrl)

	u, err := url.Parse(httpUrl)
	
	if err != nil {
		panic(err)
		return
	}

	fmt.Printf("escaped httpUrl: %v\n", u.String())

	req, err := http.NewRequest("GET", sc.ApiUrl + "/resolve.json?url=" + httpUrl + "&" + "client_id=" + sc.ClientID, nil)
	resp, err := sc.HttpClient.Do(req)

	trackid = 0
	var track *Track
	
	if err != nil {
		return trackid, err
	}
	if resp != nil {
		defer resp.Body.Close()
		
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return trackid, err
		}

		err = json.Unmarshal(bodyBytes, &track)
		
		if err != nil {
			return trackid, err
		}
		fmt.Printf("found track: %v\n", track)
		return track.Id, err
	}

	return trackid, nil
}

func (sc *Client) GetRelatedTracksByUrl(httpUrl string) (trackList []Track, err error) {
	trackid, err := sc.GetTrackIdFromUrl(httpUrl)

	if err != nil {
		panic(err)
	}

	return sc.GetRelatedTracksByID(trackid)	
}

func (sc *Client) GetRelatedTracksByID(id int) (trackList []Track, err error){

	req, err := http.NewRequest("GET", sc.ApiUrl + "/tracks/" + strconv.Itoa(id) + "/related?" + "client_id=" + sc.ClientID, nil)
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
			return strings.Contains(s.License,"cc-") && !strings.Contains(s.License,"cc-by-nc-sa")
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
