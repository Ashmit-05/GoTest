package concurrency

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
  if url == "waat://furhurterwe.geds" {
    return false
  }

  return true
}

func TestWebsites(t *testing.T) {
  websites := []string{
    "https://google.com",
    "https://github.com",
    "waat://furhurterwe.geds",
  }
  
  want := map[string]bool{
    "https://google.com" : true,
    "https://github.com" : true,
    "waat://furhurterwe.geds" : false,
  }

  got := CheckWebsites(mockWebsiteChecker,websites)

  if !reflect.DeepEqual(want,got) {
    t.Fatalf("wanted %v, got %v",want,got)
  }
}

func TestRacer(t *testing.T) {

	slowServer := makeDelayedServer(20 * time.Millisecond)
	fastServer := makeDelayedServer(0 * time.Millisecond)

	defer slowServer.Close()
	defer fastServer.Close()

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
