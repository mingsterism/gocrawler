package functions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"

	"github.com/mmcdole/gofeed"
)

// Check : accepts an error and returns an error. If err != nil, it returns panic(e)
func Check(e error) error {
	if e != nil {
		panic(e)
	} else {
		return e
	}
}

// GetHTTPResponse : accepts a url and returns a *http.Response
func GetHTTPResponse(u string) *http.Response {
	resp, err := http.Get(u)
	Check(err)
	return resp
}

// WriteStringToFile : accepts a filename string and writes uint8 code to the filename.
func WriteStringToFile(filename string, c []uint8) error {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
	Check(err)
	file, err := f.WriteString(string(c))
	Check(err)
	f.Sync()

	fmt.Printf("wrote %d bytes\n", file)
	return err
}

// NewsFeedStruct : A struct containing Title string and Link string
type NewsFeedStruct struct {
	Title string
	Link  string
}

// ParseXMLFile : accepts a NewsFeedStruct and file string and returns a channel []byte
func ParseXMLFile(s NewsFeedStruct, file string) chan []byte {
	dat, err := ioutil.ReadFile(file)
	Check(err)
	c := string(dat)
	fmt.Println(reflect.TypeOf(c))
	fp := gofeed.NewParser()
	feed, _ := fp.ParseString(c)
	feed1 := feed.Items
	j := make(chan []byte)
	go func() {
		for i := 0; i < len(feed1); i++ {
			s.Title = feed1[i].Title
			s.Link = feed1[i].Link
			a, err := json.MarshalIndent(s, "", " ")
			Check(err)
			j <- a
		}
		close(j)
	}()
	return j
}

// ParseXMLUrl : accepts a NewsFeedStruct and an xml url string and returns a channel []byte
func ParseXMLUrl(s NewsFeedStruct, u string) chan []byte {
	resp := GetHTTPResponse(u)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	Check(err)
	c := string(body)

	fp := gofeed.NewParser()
	feed, _ := fp.ParseString(c)
	feed1 := feed.Items
	j := make(chan []byte)
	go func() {
		for i := 0; i < len(feed1); i++ {
			s.Title = feed1[i].Title
			s.Link = feed1[i].Link
			a, err := json.MarshalIndent(s, "", " ")
			Check(err)
			j <- a
		}
		close(j)
	}()
	return j
}
