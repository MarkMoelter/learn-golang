package concurrency

import (
	"reflect"
	"testing"
	"time"
)

const invalidWebsite = "waat://furhurterwe.geds"

func mockWebsiteChecker(url string) bool {
	return url != invalidWebsite
}

func TestConcurrency(t *testing.T) {
	t.Run("test a mix of websites", func(t *testing.T) {
		websites := []string{
			"http://google.com",
			"http://blog.gypsydave5.com",
			invalidWebsite,
		}

		want := map[string]bool{
			"http://google.com":		  true,
			"http://blog.gypsydave5.com": true,
			invalidWebsite:			      false,		
		}

		got := CheckWebsites(mockWebsiteChecker, websites)

		if !reflect.DeepEqual(want, got) {
			t.Fatalf("wanted %v, got %v", want, got)
		}
	})
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

	for b.Loop() {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
