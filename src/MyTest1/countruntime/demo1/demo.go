package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	respTime = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "response_time",
			Help: "cost time per request",
		},
		[]string{"costTime"},
	)
)

func main() {
	urls := []string{"https://www.baidu.com", "https://www.baidu.com"}

	for _, url := range urls {
		request(url)
	}
}

func request(url string) {
	startTime := time.Now()
	response, _ := http.Get(url)
	defer response.Body.Close()
	_, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	costTime := time.Since(startTime)
	fmt.Printf("costtime:%d\n", costTime)

	respTime.WithLabelValues(fmt.Sprintf("%d", costTime)).Observe(costTime.Seconds())
}
