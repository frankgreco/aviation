module github.com/frankgreco/aviation/query

require (
	github.com/akrylysov/algnhsa v0.12.1
	github.com/frankgreco/aviation/api v0.0.0
	github.com/frankgreco/aviation/search v0.0.0
	github.com/go-kit/kit v0.8.0
	github.com/gorilla/mux v1.7.4
)

replace (
	github.com/frankgreco/aviation/api => ../api
	github.com/frankgreco/aviation/search => ../search
	github.com/frankgreco/aviation/utils => ../utils
)
