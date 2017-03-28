package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"

	"github.com/mingsterism/gocrawler/primefunctions/functions"
)

// type NewsFeed struct {
// 	Title string
// 	Link  string
// }

func printString(c <-chan []string, wg *sync.WaitGroup) {
	for b := range c {
		fmt.Println(b)
	}
	wg.Done()
}

func main() {
	// Example 0 ------------------------------------------------
	f, err := os.Open("feeds.txt")
	functions.Check(err)
	defer f.Close()

	wg := &sync.WaitGroup{}
	workerCount := runtime.NumCPU()
	wg.Add(workerCount)
	c := make(chan []string, 100)
	for i := 0; i < workerCount; i++ {
		go printString(c, wg)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		c <- strings.Fields(scanner.Text())
	}
	close(c)
	wg.Wait()

	// wg.Wait()
	// Example 1 ------------------------------------------------
	// nf := functions.NewsFeedStruct{}
	// a1 := functions.ParseXMLFile(nf, "guardian.xml")
	// for b := range a1 {
	// fmt.Println(string(b))
	// }

	// Example 2 ------------------------------------------------
	// nf := functions.NewsFeedStruct{}
	// a1 := functions.ParseXMLUrl(nf, "http://www.theguardian.com/uk/rss")
	// for b := range a1 {
	// 	fmt.Println(string(b))
	// }

}
