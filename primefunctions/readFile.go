package main

import (
	"encoding/xml"
	"fmt"
	"github.com/mingsterism/gocrawler/primefunctions/functions"
	"io/ioutil"
)

type Result1 struct {
	XMLName xml.Name `xml:"Person"`
	Name    string   `xml:"FullName"`
	City    string
}

type Channel struct {
	XMLName     xml.Name `xml:"channel"`
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	Description string   `xml:"description"`
}

type rss struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `"xml.channel"`
}

func walk(nodes []Node, f func(Node) bool) {
	for _, n := range nodes {
		if f(n) {
			walk(n.Nodes, f)
		}
	}
}

func main() {
	dat, err := ioutil.ReadFile("text.xml")
	functions.Check(err)
	c := string(dat)

	v := rss{}
	err = xml.Unmarshal([]byte(c), &v)
	functions.Check(err)
	fmt.Printf("XMLName: %#v\n", v.Channel.XMLName)
	fmt.Printf("Link: %q\n", v.Channel.Link)
	fmt.Printf("Title: %q\n", v.Channel.Title)
	fmt.Printf("Title: %q\n", v.Channel.Description)

}
