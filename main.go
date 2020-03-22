package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type (
	// Config as init
	Config struct {
		Token string
	}
)

// NewAhrefsAPI initialised the ahrefs api
func NewAhrefsAPI(c Config) *Config {
	return &c
}

func getURL(method, target, mode, token string) string {
	baseURL := "https://apiv2.ahrefs.com"
	outputType := "json"

	return fmt.Sprintf("%s?token=%s&from=%s&target=%s&ouput=%s&mode=%s", baseURL, token, method, target, outputType, mode)
}

func ahrefsRank(target, mode, token string) string {
	return getURL("ahrefs_rank", target, mode, token)
}

func anchors(target, mode, token string) string {
	return getURL("anchors", target, mode, token)
}

func anchorsRefdomains(target, mode, token string) string {
	return getURL("anchors_refdomains", target, mode, token)
}

func backlinks(target, mode, token string) string {
	return getURL("backlinks", target, mode, token)
}

func backlinksNewLost(target, mode, token string) string {
	return getURL("backlinks_new_lost", target, mode, token)
}

func backlinksNewLostCounters(target, mode, token string) string {
	return getURL("backlinks_new_lost_counters", target, mode, token)
}

func backlinksOnePerDomain(target, mode, token string) string {
	return getURL("backlinks_one_per_domain", target, mode, token)
}

func brokenBacklinks(target, mode, token string) string {
	return getURL("broken_backlinks", target, mode, token)
}

func brokenLinks(target, mode, token string) string {
	return getURL("broken_links", target, mode, token)
}

func domainRating(target, mode, token string) string {
	return getURL("domain_rating", target, mode, token)
}

func linkedAnchors(target, mode, token string) string {
	return getURL("linked_anchors", target, mode, token)
}

func linkedDomains(target, mode, token string) string {
	return getURL("linked_domains", target, mode, token)
}

func linkedDomainsByType(target, mode, token string) string {
	return getURL("linked_domains_by_type", target, mode, token)
}

func metrics(target, mode, token string) string {
	return getURL("metrics", target, mode, token)
}

func metricsExtended(target, mode, token string) string {
	return getURL("metrics_extended", target, mode, token)
}

func pages(target, mode, token string) string {
	return getURL("pages", target, mode, token)
}

func pagesExtended(target, mode, token string) string {
	return getURL("pages_extended", target, mode, token)
}

func pagesInfo(target, mode, token string) string {
	return getURL("pages_info", target, mode, token)
}

func refdomains(target, mode, token string) string {
	return getURL("refdomains", target, mode, token)
}

func refdomainsByType(target, mode, token string) string {
	return getURL("refdomains_by_type", target, mode, token)
}

func refdomainsNewLost(target, mode, token string) string {
	return getURL("refdomains_new_lost", target, mode, token)
}

func refdomainsNewLostCounters(target, mode, token string) string {
	return getURL("refdomains_new_lost_counters", target, mode, token)
}

func refips(target, mode, token string) string {
	return getURL("refips", target, mode, token)
}

func subscriptionInfo(target, mode, token string) string {
	return getURL("subscription_info", target, mode, token)

}

/* Request allows you to get decoded data from the api as Json format using stdlib http get */
func request(req string) string {
	resp, err := http.Get(req)

	if err != nil {
		log.Fatal("Error calling the page")
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(responseData)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	api := NewAhrefsAPI(Config{Token: os.Getenv("AHREFS_TOKEN")})

	fmt.Println(request(ahrefsRank("ahrefs.com", "domain", *&api.Token)))
}
