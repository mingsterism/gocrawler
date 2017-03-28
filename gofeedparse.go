package main

import (
	"fmt"
	"reflect"

	"github.com/mingsterism/gocrawler/primefunctions/functions"
)

type NewsFeed struct {
	Title string
	Link  string
}

func main() {
	// resp := functions.GetHTTPResponse("https://www.theguardian.com/world/rss")
	// defer resp.Body.Close()

	// body, err := ioutil.ReadAll(resp.Body)
	// functions.Check(err)

	// c := string(body)
	// functions.WriteStringToFile("guardian.rss", body)

	// dat, err := ioutil.ReadFile("guardian.xml")
	// functions.Check(err)
	// c := string(dat)
	// fmt.Println(reflect.TypeOf(c))
	// fp := gofeed.NewParser()
	// feed, _ := fp.ParseString(c)
	// feed1 := feed.Items
	// group := NewsFeed{} // a struct with Title and Link fields
	// for i := 0; i < len(feed1); i++ {
	// 	group.Title = feed1[i].Title
	// 	group.Link = feed1[i].Link
	// 	// group = NewsFeed{
	// 	// 	Title: feed1[i].Title,
	// 	// 	Link:  feed1[i].Link,
	// 	// }
	// 	b, err := json.MarshalIndent(group, "", " ")
	// 	functions.Check(err)
	// 	os.Stdout.Write(b)
	// }

	nf := NewsFeed{}
	a1 := functions.ParseXMLFeed(nf)
	fmt.Println(reflect.TypeOf(nf))
	fmt.Println(a1)
}
