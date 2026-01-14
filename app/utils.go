package main

import (
	"math/rand/v2"
	"net/url"
	"regexp"
)

func generateID(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)

	for i := range b {
		b[i] = letters[rand.IntN(len(letters))]
	}
	return string(b)
}

var domainRegex = regexp.MustCompile(`^([a-zA-Z0-9-]+\.)+[a-zA-Z]{2,}$`)

func isValidURL(rawURL string) bool {
	u, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return false
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		return false
	}

	if u.Host == "" {
		return false
	}

	host := u.Hostname()

	if !domainRegex.MatchString(host) {
		if host == "localhost" {
			return true
		}
		return false
	}

	return true
}
