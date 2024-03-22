// This file was auto-generated by Fern from our API Definition.

package client

import (
	core "github.com/minaelee/ferntest-go-sdk/core"
	option "github.com/minaelee/ferntest-go-sdk/option"
	pet "github.com/minaelee/ferntest-go-sdk/pet"
	store "github.com/minaelee/ferntest-go-sdk/store"
	user "github.com/minaelee/ferntest-go-sdk/user"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header

	Pet   *pet.Client
	Store *store.Client
	User  *user.Client
}

func NewClient(opts ...option.RequestOption) *Client {
	options := core.NewRequestOptions(opts...)
	return &Client{
		baseURL: options.BaseURL,
		caller: core.NewCaller(
			&core.CallerParams{
				Client:      options.HTTPClient,
				MaxAttempts: options.MaxAttempts,
			},
		),
		header: options.ToHeader(),
		Pet:    pet.NewClient(opts...),
		Store:  store.NewClient(opts...),
		User:   user.NewClient(opts...),
	}
}
