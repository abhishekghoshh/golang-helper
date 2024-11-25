module github.com/abhishekghoshh/rate-limiting

go 1.22.0

require golang.org/x/time v0.8.0

require github.com/patrickmn/go-cache v2.1.0+incompatible // indirect

require (
	github.com/didip/tollbooth v4.0.2+incompatible
	github.com/didip/tollbooth/v8 v8.0.0 // indirect
	github.com/go-pkgz/expirable-cache/v3 v3.0.0 // indirect
)
