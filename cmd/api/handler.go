package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	httptransport "github.com/go-kit/kit/transport/http"
)

func makeHandler(s Service) http.Handler {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(encodeError),
	}

	addHandler := httptransport.NewServer(
		makeAddEnpoint(s),
		decodeAddRequest,
		encodeAddResponse,
		options...,
	)

	removeHandler := httptransport.NewServer(
		makeRemoveEndpoint(s),
		decodeRemoveRequest,
		encodeRemoveResponse,
		options...,
	)

	getAllHandler := httptransport.NewServer(
		makeGetAllEndpoint(s),
		decodeGetAllRequest,
		encodeGetAllResponse,
		options...,
	)

	r := chi.NewRouter()
	r.Route("/items", func(chi.Router) {
		r.Get("/", getAllHandler.ServeHTTP)
		r.Post("/add", addHandler.ServeHTTP)
		r.Get("/remove/{ID}", removeHandler.ServeHTTP)
	})

	return r
}
