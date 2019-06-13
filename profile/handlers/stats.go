package handlers

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/varstr/uaparser"
)

var printStats = flag.Bool("printStats", false, "Print stats to console")

// WithStats wraps handlers with stats reporting. It tracks metrics such
// as the number of requests per endpoint, the latency, etc.
func WithStats(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		tags := getStatsTags(r)
		incCounter("handler.received", tags, 1)

		h(w, r)

		duration := time.Since(start)
		recordTimer("handler.latency", tags, duration)
	}
}

func getStatsTags(r *http.Request) map[string]string {
	userBrowser, userOS := parseUserAgent(r.UserAgent())
	stats := map[string]string{
		"browser":  userBrowser,
		"os":       userOS,
		"endpoint": filepath.Base(r.URL.Path),
	}
	host, err := os.Hostname()
	if err == nil {
		if idx := strings.IndexByte(host, '.'); idx > 0 {
			host = host[:idx]
		}
		stats["host"] = host
	}
	return stats
}

func parseUserAgent(uaString string) (browser, os string) {
	ua := uaparser.Parse(uaString)

	if ua.Browser != nil {
		browser = ua.Browser.Name
	}
	if ua.OS != nil {
		os = ua.OS.Name
	}

	return browser, os
}

// incCounter increments a counter.
func incCounter(name string, tags map[string]string, value int64) {
	name = addTagsToName(name, tags)
	if *printStats {
		log.Printf("IncCounter: %v = %v\n", name, value)
	}
}

// recordTimer records a timer.
func recordTimer(name string, tags map[string]string, d time.Duration) {
	name = addTagsToName(name, tags)
	if *printStats {
		log.Printf("RecordTimer: %v = %v\n", name, d)
	}
}

func addTagsToName(name string, tags map[string]string) string {
	// The format we want is: host.endpoint.os.browser
	// if there's no host tag, then we don't use it.
	var keyOrder []string
	if _, ok := tags["host"]; ok {
		keyOrder = append(keyOrder, "host")
	}
	keyOrder = append(keyOrder, "endpoint", "os", "browser")

	parts := []string{name}
	for _, k := range keyOrder {
		v, ok := tags[k]
		if !ok || v == "" {
			parts = append(parts, "no-"+k)
			continue
		}
		parts = append(parts, clean(v))
	}

	return strings.Join(parts, ".")
}

var specialChars = regexp.MustCompile(`[{}/\\:\s.]`)

// clean takes a string that may contain special characters, and replaces these
// characters with a '-'.
func clean(value string) string {
	return specialChars.ReplaceAllString(value, "-")
}
