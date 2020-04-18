package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/joho/godotenv"
)

type (
	// Config as init
	Config struct {
		Token string
		Ouput string
	}

	// Request for ahrefs data quries
	Request struct {
		Target  string `url:"target"`
		Mode    string `url:"mode"`
		Limit   int    `url:"limit,omitempty"`
		OrderBy string `url:"order_by,omitempty"`
		where   string `url:"where,omitempty"`
		having  string `url:"having,omitempty"`
	}

	// AhrefsRankList - AhrefsRank -> Page
	AhrefsRankList struct {
		URL string `json:"url"`
		AR  int    `json:"ahrefs_rank"`
	}

	// AhrefsRank - Contains the URLs and their rankings.
	AhrefsRank struct {
		Pages []AhrefsRankList
	}

	// AnchorsStats - Anchors -> States
	AnchorsStats struct {
		Backlinks int
		Refpages  int
	}

	// AnchorsAnchors - Anchors -> Anchors
	AnchorsAnchors struct {
		Anchor      string
		Backlinks   int
		Refpages    int
		Refdomains  int
		FirstSeen   time.Time `json:"first_seen"`
		LastVisited time.Time `json:"last_visited"`
	}

	// Anchors - Contains the anchor text and the number of backlinks, referring pages and referring domains that has it.
	Anchors struct {
		Anchors []AnchorsAnchors
		Stats   AnchorsStats
	}

	// AnchorsRefdomainsRefdomains - AnchorsRefdomains -> Refdomains
	AnchorsRefdomainsRefdomains struct {
		Anchor     string
		Backlinks  int
		Refdomains int
	}

	// AnchorsRefdomains - Contains connection between anchors and domains. Can be used to get all referring domains with specified anchor.
	AnchorsRefdomains struct {
		Refdomains []AnchorsRefdomainsRefdomains
	}
)

// NewAhrefsAPI initialised the ahrefs api
func NewAhrefsAPI(token string) Config {
	return Config{Token: token, Ouput: "json"}
}

func getURL(method string, r Request, c *Config) string {
	baseURL := "https://apiv2.ahrefs.com"

	query, _ := query.Values(r)

	return fmt.Sprintf("%s?token=%s&from=%s&%s", baseURL, c.Token, method, query.Encode())
}

/* Request allows you to get decoded data from the api as Json format using stdlib http get */
func request(req string) []byte {
	resp, err := http.Get(req)

	if err != nil {
		log.Fatal("Error calling the page")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return body
}

func ahrefsRank(r Request, c *Config) *AhrefsRank {
	responseData := request(getURL("ahrefs_rank", r, c))

	ahrefsRank := &AhrefsRank{}
	decoder := json.NewDecoder(bytes.NewReader(responseData))

	err := decoder.Decode(ahrefsRank)
	if err != nil {
		log.Fatal(err)
	}

	return ahrefsRank
}

func anchors(r Request, c *Config) *Anchors {
	responseData := request(getURL("anchors", r, c))

	anchors := &Anchors{}
	decoder := json.NewDecoder(bytes.NewReader(responseData))

	err := decoder.Decode(anchors)
	if err != nil {
		log.Fatal(err)
	}

	return anchors
}

func anchorsRefdomains(r Request, c *Config) *AnchorsRefdomains {
	responseData := request(getURL("anchors_refdomains", r, c))

	anchorsRefdomains := &AnchorsRefdomains{}
	decoder := json.NewDecoder(bytes.NewReader(responseData))

	err := decoder.Decode(anchorsRefdomains)
	if err != nil {
		log.Fatal(err)
	}

	return anchorsRefdomains
}

func backlinks(r Request, c *Config) string {
	return getURL("backlinks", r, c)
}

func backlinksNewLost(r Request, c *Config) string {
	return getURL("backlinks_new_lost", r, c)
}

func backlinksNewLostCounters(r Request, c *Config) string {
	return getURL("backlinks_new_lost_counters", r, c)
}

func backlinksOnePerDomain(r Request, c *Config) string {
	return getURL("backlinks_one_per_domain", r, c)
}

func brokenBacklinks(r Request, c *Config) string {
	return getURL("broken_backlinks", r, c)
}

func brokenLinks(r Request, c *Config) string {
	return getURL("broken_links", r, c)
}

func domainRating(r Request, c *Config) string {
	return getURL("domain_rating", r, c)
}

func linkedAnchors(r Request, c *Config) string {
	return getURL("linked_anchors", r, c)
}

func linkedDomains(r Request, c *Config) string {
	return getURL("linked_domains", r, c)
}

func linkedDomainsByType(r Request, c *Config) string {
	return getURL("linked_domains_by_type", r, c)
}

func metrics(r Request, c *Config) string {
	return getURL("metrics", r, c)
}

func metricsExtended(r Request, c *Config) string {
	return getURL("metrics_extended", r, c)
}

func pages(r Request, c *Config) string {
	return getURL("pages", r, c)
}

func pagesExtended(r Request, c *Config) string {
	return getURL("pages_extended", r, c)
}

func pagesInfo(r Request, c *Config) string {
	return getURL("pages_info", r, c)
}

func refdomains(r Request, c *Config) string {
	return getURL("refdomains", r, c)
}

func refdomainsByType(r Request, c *Config) string {
	return getURL("refdomains_by_type", r, c)
}

func refdomainsNewLost(r Request, c *Config) string {
	return getURL("refdomains_new_lost", r, c)
}

func refdomainsNewLostCounters(r Request, c *Config) string {
	return getURL("refdomains_new_lost_counters", r, c)
}

func refips(r Request, c *Config) string {
	return getURL("refips", r, c)
}

func subscriptionInfo(r Request, c *Config) string {
	return getURL("subscription_info", r, c)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := NewAhrefsAPI(os.Getenv("AHREFS_TOKEN"))

	fmt.Println(ahrefsRank(Request{Target: "ahrefs.com", Mode: "domain"}, &config))
	fmt.Println(anchors(Request{Target: "ahrefs.com", Mode: "domain"}, &config))
}
