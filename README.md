# ahrefs api go wrapper

This repo is for learning purpose, learning golang by creating a ahrefs api go wrapper.

```
err := godotenv.Load()
if err != nil {
  log.Fatal("Error loading .env file")
}

api := NewAhrefsAPI(Config{Token: os.Getenv("AHREFS_TOKEN")})

fmt.Println(request(ahrefsRank("ahrefs.com", "domain", *&api.Token)))

```