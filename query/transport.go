package main

import (
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"github.com/frankgreco/aviation/api"
)

// MakeHTTPHandler mounts all of the service endpoints into an http.Handler.
func MakeHTTPHandler(s Service) http.Handler {
	r := mux.NewRouter()
	e := MakeServerEndpoints(s)

	r.Methods(http.MethodGet).Path("/search").Handler(httptransport.NewServer(
		e.SearchEndpoint,
		decodeSearchRequest,
		encodeSearchResponse,
	))

	return r
}

func encodeSearchResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(api.Errorer); ok && e.GetError() != nil {
		w.WriteHeader(int(e.GetError().Code))
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func decodeSearchRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	return api.SearchRequest{
		Query: r.URL.Query().Get("q"),
	}, nil
}
