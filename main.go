package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Data struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

type Status struct {
	Water string `json:"Status Water"`
	Wind  string `json:"Status Wind"`
}

func main() {
	http.HandleFunc("/post", postData)
	http.ListenAndServe(":8080", nil)
}

func postData(w http.ResponseWriter, r *http.Request) {
	interval := 15 * time.Second

	for {
		rand.Seed(time.Now().UnixNano())

		url := "https://jsonplaceholder.typicode.com/posts"

		valueWater := rand.Intn(100) + 1
		valueWind := rand.Intn(100) + 1

		data := map[string]interface{}{
			"water": valueWater,
			"wind":  valueWind,
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
		req.Header.Set("X-Custom-Header", "myvalue")
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		var result Data
		err = json.Unmarshal(body, &result)
		if err != nil {
			log.Fatal(err)
		}

		statusWater := getStatusWater(result.Water)
		statusWind := getStatusWind(result.Wind)

		//strResBody := string(result)
		fmt.Printf("%+v\n", result)
		fmt.Printf("status water : %s \n", statusWater)
		fmt.Printf("status wind : %s \n", statusWind)

		time.Sleep(interval)
	}
}

func getStatusWater(value int) string {
	if value < 5 {
		return "aman"
	} else if value >= 5 && value <= 8 {
		return "siaga"
	} else {
		return "bahaya"
	}
}

func getStatusWind(value int) string {
	if value < 6 {
		return "aman"
	} else if value >= 6 && value <= 15 {
		return "siaga"
	} else {
		return "bahaya"
	}
}
