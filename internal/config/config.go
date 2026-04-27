package config

import "time"

type Config struct {
	Timeout time.Duration
	Sites   []string
}

func Load() *Config {
	return &Config{
		Timeout: 5 * time.Second,
		Sites: []string{
			"https://www.google.com",
			"https://www.youtube.com",
			"https://www.facebook.com",
			"https://www.instagram.com",
			"https://www.chatgpt.com",
			"https://www.x.com",
			"https://www.wikipedia.org",
			"https://www.reddit.com",
			"https://www.yahoo.com",
			"https://www.amazon.com",
		},
	}
}
