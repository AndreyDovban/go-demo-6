package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	t := time.Now()
	var wg sync.WaitGroup

	for i := range 10 {
		wg.Add(1)
		go getGoogle(&wg, i)
	}

	wg.Wait()

	fmt.Println(time.Since(t))

}

func getGoogle(wg *sync.WaitGroup, i int) {

	defer wg.Done()

	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Print(i, " : ")
	fmt.Println(resp.StatusCode)

}
