package concurrency

import (
	"net/http"
	"time"
)

// import "time"

type WebsiteChecker func(string) bool

type result struct {
  string
  bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
  results := make(map[string]bool)
  resultChannel := make(chan result)

  for _,url := range urls {
    go func (u string)  { // this is an anonymous function call
      resultChannel <- result{u,wc(u)}
    }(url)
  }

  for i := 0; i < len(urls); i++ {
    r := <-resultChannel
    results[r.string] = r.bool
  }

  // time.Sleep(2 * time.Second)

  return results
}

func Racer(a, b string) (winner string) {
  aDuration := measureResponseTime(a)
  bDuration := measureResponseTime(b)

	if aDuration < bDuration {
		return a
	}

	return b
}

func measureResponseTime(a string) time.Duration {
  aDuration := time.Now()
  http.Get(a)
  startA := time.Since(aDuration)
  return startA
}
