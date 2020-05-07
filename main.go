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
		Backlinks int `json:"backlinks"`
		Refpages  int `json:"refpages"`
	}

	// AnchorsAnchors - Anchors -> Anchors
	AnchorsAnchors struct {
		Anchor      string    `json:"anchor"`
		Backlinks   int       `json:"backlinks"`
		Refpages    int       `json:"refpages"`
		Refdomains  int       `json:"refdomains"`
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
		Anchor     string `json:"anchor"`
		Backlinks  int    `json:"backlinks"`
		Refdomains int    `json:"refdomains"`
	}

	// AnchorsRefdomains - Contains connection between anchors and domains. Can be used to get all referring domains with specified anchor.
	AnchorsRefdomains struct {
		Refdomains []AnchorsRefdomainsRefdomains
	}

	// BacklinksRefPages - Backlinks - RefPages
	BacklinksRefPages struct {
		Date             time.Time
		Type             string
		From             string    `json:"url_from"`
		To               string    `json:"url_to"`
		AR               int       `json:"ahrefs_rank"`
		DomainRating     int       `json:"domain_rating"`
		AhrefsTop        int       `json:"ahrefs_top"`
		IPFrom           string    `json:"ip_from"`
		LinksInternal    int       `json:"links_internal"`
		LinksExternal    int       `json:"links_external"`
		PageSize         int       `json:"page_size"`
		Encoding         string    `json:"encoding"`
		Language         string    `json:"language"`
		Title            string    `json:"title"`
		FirstSeen        time.Time `json:"first_seen"`
		LastVisited      time.Time `json:"last_visited"`
		PrevVisited      time.Time `json:"prev_visited"`
		Original         bool      `json:"original"`
		LinkType         string    `json:"link_type"`
		Redirect         int       `json:"redirect"`
		NoFollow         bool      `json:"nofollow"`
		Alt              string    `json:"alt"`
		Anchor           string    `json:"anchor"`
		TextPrev         string    `json:"text_pre"`
		TextPost         string    `json:"text_post"`
		HttoCode         int       `json:"http_code"`
		URLFromFirstSeen string    `json:"url_from_first_seen"` // time.Time type parsing error in here...
	}

	// Backlinks - Contains the backlinks and details of the referring pages, such as anchor and page title.
	Backlinks struct {
		RefPages []BacklinksRefPages
	}

	// SubscriptionInfo - Contains user subscription information.
	SubscriptionInfo struct {
		RowsLeft     int    `json:"rows_left"`
		RowsLimit    int    `json:"rows_limit"`
		Subscription string `json:"subscription"`
	}

	// RefipsRefdomains - Refips -> Refdomains
	RefipsRefdomains struct {
		Refip     string `json:"refip"`
		Refdomain string `json:"refdomain"`
		Backlinks int    `json:"backlinks"`
	}

	// Refips - Returns the referring IP addresses that have at least one link to the target.
	Refips struct {
		Refdomains RefipsRefdomains
	}

	// DomainRatingDomain - Refips -> Refdomains
	DomainRatingDomain struct {
		DomainRating string `json:"domain_rating"`
		AhrefsTop    int    `json:"ahrefs_top"`
	}

	// DomainRating - Contains the Domain Rating.
	// Refer to Principles of Domain Rating calculation for more information about Domain Rating.
	DomainRating struct {
		Domain DomainRatingDomain
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

func backlinks(r Request, c *Config) *Backlinks {
	responseData := request(getURL("backlinks", r, c))

	backlinks := &Backlinks{}
	decoder := json.NewDecoder(bytes.NewReader(responseData))

	err := decoder.Decode(backlinks)
	if err != nil {
		log.Fatal(err)
	}

	return backlinks
}

func backlinksNewLost(r Request, c *Config) *Backlinks {
	responseData := request(getURL("backlinks_new_lost", r, c))

	backlinks := &Backlinks{}
	decoder := json.NewDecoder(bytes.NewReader(responseData))

	err := decoder.Decode(backlinks)
	if err != nil {
		log.Fatal(err)
	}

	return backlinks
}

func backlinksNewLostCounters(r Request, c *Config) string {
	return getURL("backlinks_new_lost_counters", r, c)
}

func backlinksOnePerDomain(r Request, c *Config) *Backlinks {
	responseData := request(getURL("backlinks_one_per_domain", r, c))

	backlinks := &Backlinks{}
	decoder := json.NewDecoder(bytes.NewReader(responseData))

	err := decoder.Decode(backlinks)
	if err != nil {
		log.Fatal(err)
	}

	return backlinks
}

func brokenBacklinks(r Request, c *Config) string {
	return getURL("broken_backlinks", r, c)
}

func brokenLinks(r Request, c *Config) string {
	return getURL("broken_links", r, c)
}

func domainRating(r Request, c *Config) *DomainRating {
	responseData := request(getURL("domain_rating", r, c))

	domainRating := &DomainRating{}
	decoder := json.NewDecoder(bytes.NewReader(responseData))

	err := decoder.Decode(domainRating)
	if err != nil {
		log.Fatal(err)
	}

	return domainRating
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

func refips(r Request, c *Config) *Refips {
	responseData := request(getURL("refips", r, c))

	refips := &Refips{}
	decoder := json.NewDecoder(bytes.NewReader(responseData))

	err := decoder.Decode(refips)
	if err != nil {
		log.Fatal(err)
	}

	return refips
}

func subscriptionInfo(r Request, c *Config) *SubscriptionInfo {
	responseData := request(getURL("subscription_info", r, c))

	subscriptionInfo := &SubscriptionInfo{}
	decoder := json.NewDecoder(bytes.NewReader(responseData))

	err := decoder.Decode(subscriptionInfo)
	if err != nil {
		log.Fatal(err)
	}

	return subscriptionInfo
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := NewAhrefsAPI(os.Getenv("AHREFS_TOKEN"))

	fmt.Println(ahrefsRank(Request{Target: "ahrefs.com", Mode: "domain"}, &config))
	fmt.Println(anchors(Request{Target: "ahrefs.com", Mode: "domain"}, &config))
	fmt.Println(anchorsRefdomains(Request{Target: "ahrefs.com", Mode: "domain"}, &config))
	fmt.Println(refips(Request{Target: "ahrefs.com", Mode: "domain"}, &config))
	fmt.Println(domainRating(Request{Target: "ahrefs.com", Mode: "domain"}, &config))
	fmt.Println(backlinks(Request{Target: "ahrefs.com", Mode: "domain"}, &config))
	fmt.Println(backlinksNewLost(Request{Target: "ahrefs.com", Mode: "domain"}, &config))
	fmt.Println(backlinksOnePerDomain(Request{Target: "ahrefs.com", Mode: "domain"}, &config))
}
