package checker

import (
	"net/http"
	"sync"
	"time"

	"go-wg/internal/entity"
)

type Checker struct {
	client *http.Client
}

func New(timeout time.Duration) *Checker {
	return &Checker{
		client: &http.Client{Timeout: timeout},
	}
}

func (c *Checker) CheckAll(urls []string) []entity.Site {
	var wg sync.WaitGroup
	results := make([]entity.Site, len(urls))

	for i, url := range urls {
		wg.Add(1)
		go func(idx int, u string) {
			defer wg.Done()
			results[idx] = c.check(u)
		}(i, url)
	}

	wg.Wait()
	return results
}

func (c *Checker) check(url string) entity.Site {
	start := time.Now()
	resp, err := c.client.Get(url)
	latency := time.Since(start)

	if err != nil {
		return entity.Site{URL: url, Latency: latency, Err: err}
	}
	defer resp.Body.Close()

	return entity.Site{
		URL:     url,
		Status:  resp.StatusCode,
		Latency: latency,
	}
}
