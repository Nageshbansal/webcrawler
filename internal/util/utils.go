package util

import (
	"github.com/Nageshbansal/web-crawler/internal/logger"
	"net/url"
	"strings"
)

func IsSameDomain(link string, domain string) bool {
	// u.strin = https://redhat.com/abc/def/
	// d.path = https://redhat.com/
	u, err := url.Parse(link)
	d, err := url.Parse(domain)
	if err != nil {
		logger.Errorf("Error parsing URL: %v", err)
		return false
	}
	// redhat.com
	return d.Hostname() == u.Hostname() && strings.Contains(u.String(), d.Path)
}

func ResolveURL(href, base string) string {
	u, err := url.Parse(href)
	if err != nil {
		return ""
	}
	baseURL, err := url.Parse(base)
	if err != nil {
		return ""
	}
	return baseURL.ResolveReference(u).String()
}
