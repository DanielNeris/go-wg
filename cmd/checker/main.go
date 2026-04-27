package main

import (
	"fmt"
	"time"

	"go-wg/internal/checker"
	"go-wg/internal/config"
	"go-wg/internal/reporter"
)

func main() {
	cfg := config.Load()

	fmt.Println("Starting processing")
	start := time.Now()

	c := checker.New(cfg.Timeout)
	results := c.CheckAll(cfg.Sites)

	reporter.Print(results)

	fmt.Printf("Processing finished. Elapsed: %.2fs\n", time.Since(start).Seconds())
}
