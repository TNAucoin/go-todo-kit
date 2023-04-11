package main

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type addRequest struct {
	Name string "json:`name`"
}

type removeRequest struct {
	ID int "json:`id`"
}

type addResponse struct {
	payload model
	err     error
}

type removeResponse struct {
	err error
}

type getAllResponse struct {
	payload []model
	err     error
}

func makeAddEnpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		input := request.(addRequest)
		payload, err := s.add(input.Name)
		return &addResponse{
			payload: payload,
			err:     err,
		}, nil
	}
}

func makeRemoveEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		input := request.(removeRequest)
		err := s.remove(input.ID)
		return &removeResponse{
			err: err,
		}, nil
	}
}

func makeGetAllEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := s.getall()
		return &getAllResponse{
			payload: res,
			err:     err,
		}, nil
	}
}
