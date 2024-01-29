package concurrency

import (
	"fmt"
	"net/http"
	"time"
)

// import "time"

type WebsiteChecker func(string) bool

type result struct {
  string
  bool
}

var tenSecondTimeout = 10 * time.Second

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

// func Racer(a, b string) (winner string) {
//   aDuration := measureResponseTime(a)
//   bDuration := measureResponseTime(b)
//
// 	if aDuration < bDuration {
// 		return a
// 	}
//
// 	return b
// }

func measureResponseTime(a string) time.Duration {
  aDuration := time.Now()
  http.Get(a)
  startA := time.Since(aDuration)
  return startA
}

/*
You'll recall from the concurrency chapter that you can wait for values to be sent to a channel with myVar := <-ch. This is a blocking call, as you're waiting for a value.
select allows you to wait on multiple channels. The first one to send a value "wins" and the code underneath the case is executed.
We use ping in our select to set up two channels, one for each of our URLs. Whichever one writes to its channel first will have its code executed in the select, which results in its URL being returned (and being the winner).
*/

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
  ch := make(chan struct{})
  go func() {
    http.Get(url)
    close(ch)
  } ()
  return ch
}
