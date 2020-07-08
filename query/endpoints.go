package main

import (
	"context"
	"encoding/base64"

	"github.com/go-kit/kit/endpoint"

	"github.com/frankgreco/aviation/api"
)

// Endpoints collects all of the endpoints that compose this service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	SearchEndpoint endpoint.Endpoint
}

// MakeServerEndpoints returns an Endpoints struct where each endpoint invokes
// the corresponding method on the provided service.
func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		SearchEndpoint: MakeSearchEndpoint(s),
	}
}

// MakeSearchEndpoint returns an endpoint via the passed service.
func MakeSearchEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(api.SearchRequest)
		resp := new(api.SearchResponse)

		query, err := base64.URLEncoding.DecodeString(req.Query)
		if err != nil {
			resp.Error = api.WrapErr(err, "could not base64 decode query").(*api.Error)
		}

		results, err := s.Search(ctx, string(query), req.Limit)
		if err != nil {
			resp.Error = err.(*api.Error)
			return resp, nil
		}
		resp.Items = results
		return resp, nil
	}
}
