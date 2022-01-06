package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	file, err := os.Create("./output.csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	for i := 1; i <= 100; i++ {
		duration := requestWithTimer(i)
		fmt.Println(duration)
		file.WriteString(fmt.Sprintf("%d,%f\n", i, duration))
		time.Sleep(time.Second * 10)
	}
	// fmt.Println(requestWithTimer(1))
}

func requestWithTimer(perPage int) float32 {
	start := time.Now()
	path := fmt.Sprintf("https://content-api.cloud.unity3d.com/api/v1/projects/bcd835c4-d711-4bed-aa17-6c1458bafb73/buckets/?page=1&per_page=%d", perPage)
	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("Authorization", "Basic <API KEY>")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(resBody))
	duration := time.Since(start)
	defer res.Body.Close()
	return float32(duration.Seconds())
}
