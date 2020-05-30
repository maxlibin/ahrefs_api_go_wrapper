# ahrefs api go wrapper

This repo is for learning purpose, learning golang by creating a ahrefs api go wrapper.

```
err := godotenv.Load()
if err != nil {
  log.Fatal("Error loading .env file")
}

api := NewAhrefsAPI(Config{Token: os.Getenv("AHREFS_TOKEN")})

fmt.Println(ahrefsRank(Request{Target: "ahrefs.com", Mode: "domain"}, &config))
fmt.Println(anchors(Request{Target: "ahrefs.com", Mode: "domain"}, &config))
fmt.Println(anchorsRefdomains(Request{Target: "ahrefs.com", Mode: "domain"}, &config))
fmt.Println(refips(Request{Target: "ahrefs.com", Mode: "domain"}, &config))
fmt.Println(domainRating(Request{Target: "ahrefs.com", Mode: "domain"}, &config))
fmt.Println(backlinks(Request{Target: "ahrefs.com", Mode: "domain"}, &config))
fmt.Println(backlinksNewLost(Request{Target: "ahrefs.com", Mode: "domain"}, &config))
fmt.Println(backlinksOnePerDomain(Request{Target: "ahrefs.com", Mode: "domain"}, &config))
fmt.Println(backlinksNewLostCounters(request, &config))
fmt.Println(backlinksNewLostCounters(request, &config))
fmt.Println(backlinksNewLostCounters(request, &config))
fmt.Println(brokenBacklinks(request, &config))
fmt.Println(brokenLinks(request, &config))
```
