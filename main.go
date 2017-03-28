package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
	"sync"
)

func main() {
	f, err := os.Open("feeds.txt")
	if err != nil {
		log.Fatal()
	}
	defer f.Close()

	wg := &sync.WaitGroup{}
	workerCount := runtime.NumCPU()
	wg.Add(workerCount)
	c := make(chan []string, 100)
	for i := 0; i < workerCount; i++ {
		go crawler(c, wg)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		c <- strings.Fields(scanner.Text())
	}
	close(c)

	wg.Wait()

}

func crawler(in <-chan []string, wg *sync.WaitGroup) {
	for url := range in {
		fmt.Printf("%#v\n", url)
		resp, err := http.Get(url[0])
		if err != nil {
			fmt.Println("there is an error +++++++++++++++++++++++++++++++++++++++++")
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
	}
	wg.Done()
}
