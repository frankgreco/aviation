module github.com/frankgreco/aviation

go 1.13

require (
	github.com/frankgreco/aviation/download v0.0.0
	github.com/frankgreco/aviation/load v0.0.0
)

replace (
	github.com/frankgreco/aviation => ./
	github.com/frankgreco/aviation/api => ./api
	github.com/frankgreco/aviation/download => ./download
	github.com/frankgreco/aviation/load => ./load
	github.com/frankgreco/aviation/utils => ./utils
)
