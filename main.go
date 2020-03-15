package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// Config as init
type Config struct {
	Token  string string `json ", omitempty", c`
}

// Options types for ahrefs
type Options struct {
	Target string `json ", omitempty"`
	Mode string `json ", omitempty"`
	Where string `json ", omitempty"`
	Having string `json ", omitempty"`
	OrderBy string `json "order_by, omitempty"`
	Limit string `json ", omitempty"`
}

// NewAhrefsAPI initialised the ahrefs api
func NewAhrefsAPI(c Config) *Config {
	return &c
}

func getURL(method, target, mode string, c Config) string {
	baseURL := "https://apiv2.ahrefs.com"

	return fmt.Sprintf("%s?from=%s&target=%s&mode=%s", baseURL, method, target, mode)
}

func ahrefsRank(target, mode string, c Config) string {
	return getURL("ahrefs_rank", target, mode, c)
}

func anchors(target, mode string, c Config):
	return getURL("anchors", target, mode, c)

func anchorsRefdomains(target, mode string, c Config):
	return getURL("anchors_refdomains", target, mode, c)

func backlinks(target, mode string, c Config):
	return getURL("backlinks", target, mode, c)

func backlinksNewLost(target, mode string, c Config):
	return getURL("backlinks_new_lost", target, mode, c)

func backlinksNewLost_counters(target, mode string, c Config):
	return getURL("backlinks_new_lost_counters", target, mode, c)

func backlinksOnePer_domain(target, mode string, c Config):
	return getURL("backlinks_one_per_domain", target, mode, c)

func brokenBacklinks(target, mode string, c Config):
	return getURL("broken_backlinks", target, mode, c)

func brokenLinks(target, mode string, c Config):
	return getURL("broken_links", target, mode, c)

func domainRating(target, mode string, c Config):
	return getURL("domain_rating", target, mode, c)

func linkedAnchors(target, mode string, c Config):
	return getURL("linked_anchors", target, mode, c)

func linkedDomains(target, mode string, c Config):
	return getURL("linked_domains", target, mode, c)

func linkedDomains_by_type(target, mode string, c Config):
	return getURL("linked_domains_by_type", target, mode, c)

func metrics(target, mode string, c Config):
	return getURL("metrics", target, mode, c)

func metrics_extended(target, mode string, c Config):
	return getURL("metrics_extended", target, mode, c)

func pages(target, mode string, c Config):
	return getURL("pages", target, mode, c)

func pagesExtended(target, mode string, c Config):
	return getURL("pages_extended", target, mode, c)

func pagesInfo(target, mode string, c Config):
	return getURL("pages_info", target, mode, c)

func refdomains(target, mode string, c Config):
	return getURL("refdomains", target, mode, c)

func refdomainsByType(target, mode string, c Config):
	return getURL("refdomains_by_type", target, mode, c)

func refdomainsNewLost(target, mode string, c Config):
	return getURL("refdomains_new_lost", target, mode, c)

func refdomainsNewLostCounters(target, mode string, c Config):
	return getURL("refdomains_new_lost_counters", target, mode, c)

func refips(target, mode string, c Config):
	return getURL("refips", target, mode, c)

func subscriptionInfo(target, mode string, c Config):
	return getURL("subscription_info", target, mode, c)

func main() {
	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
	}

	api := NewAhrefsAPI(Config{Token: os.Getenv("AHREFS_TOKEN")})

	fmt.Println(ahrefsRank("ahrefs.com", "domain", *api))
}
