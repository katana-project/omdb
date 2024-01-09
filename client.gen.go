// Package omdb provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.0 DO NOT EDIT.
package omdb

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/oapi-codegen/runtime"
)

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetId request
	GetId(ctx context.Context, params *GetIdParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// TitleSearch request
	TitleSearch(ctx context.Context, params *TitleSearchParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetTitle request
	GetTitle(ctx context.Context, params *GetTitleParams, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetId(ctx context.Context, params *GetIdParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetIdRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) TitleSearch(ctx context.Context, params *TitleSearchParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewTitleSearchRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetTitle(ctx context.Context, params *GetTitleParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetTitleRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetIdRequest generates requests for GetId
func NewGetIdRequest(server string, params *GetIdParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/?i")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "i", runtime.ParamLocationQuery, params.I); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

		if params.Plot != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "plot", runtime.ParamLocationQuery, *params.Plot); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.R != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "r", runtime.ParamLocationQuery, *params.R); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.Callback != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "callback", runtime.ParamLocationQuery, *params.Callback); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewTitleSearchRequest generates requests for TitleSearch
func NewTitleSearchRequest(server string, params *TitleSearchParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/?s")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "s", runtime.ParamLocationQuery, params.S); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

		if params.Y != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "y", runtime.ParamLocationQuery, *params.Y); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.Type != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "type", runtime.ParamLocationQuery, *params.Type); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.R != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "r", runtime.ParamLocationQuery, *params.R); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.Page != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "page", runtime.ParamLocationQuery, *params.Page); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.Callback != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "callback", runtime.ParamLocationQuery, *params.Callback); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetTitleRequest generates requests for GetTitle
func NewGetTitleRequest(server string, params *GetTitleParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/?t")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "t", runtime.ParamLocationQuery, params.T); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

		if params.Y != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "y", runtime.ParamLocationQuery, *params.Y); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.Type != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "type", runtime.ParamLocationQuery, *params.Type); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.Plot != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "plot", runtime.ParamLocationQuery, *params.Plot); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.R != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "r", runtime.ParamLocationQuery, *params.R); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.Callback != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "callback", runtime.ParamLocationQuery, *params.Callback); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// GetIdWithResponse request
	GetIdWithResponse(ctx context.Context, params *GetIdParams, reqEditors ...RequestEditorFn) (*GetIdResponse, error)

	// TitleSearchWithResponse request
	TitleSearchWithResponse(ctx context.Context, params *TitleSearchParams, reqEditors ...RequestEditorFn) (*TitleSearchResponse, error)

	// GetTitleWithResponse request
	GetTitleWithResponse(ctx context.Context, params *GetTitleParams, reqEditors ...RequestEditorFn) (*GetTitleResponse, error)
}

type GetIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r GetIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type TitleSearchResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r TitleSearchResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r TitleSearchResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetTitleResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r GetTitleResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetTitleResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetIdWithResponse request returning *GetIdResponse
func (c *ClientWithResponses) GetIdWithResponse(ctx context.Context, params *GetIdParams, reqEditors ...RequestEditorFn) (*GetIdResponse, error) {
	rsp, err := c.GetId(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetIdResponse(rsp)
}

// TitleSearchWithResponse request returning *TitleSearchResponse
func (c *ClientWithResponses) TitleSearchWithResponse(ctx context.Context, params *TitleSearchParams, reqEditors ...RequestEditorFn) (*TitleSearchResponse, error) {
	rsp, err := c.TitleSearch(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseTitleSearchResponse(rsp)
}

// GetTitleWithResponse request returning *GetTitleResponse
func (c *ClientWithResponses) GetTitleWithResponse(ctx context.Context, params *GetTitleParams, reqEditors ...RequestEditorFn) (*GetTitleResponse, error) {
	rsp, err := c.GetTitle(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetTitleResponse(rsp)
}

// ParseGetIdResponse parses an HTTP response from a GetIdWithResponse call
func ParseGetIdResponse(rsp *http.Response) (*GetIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseTitleSearchResponse parses an HTTP response from a TitleSearchWithResponse call
func ParseTitleSearchResponse(rsp *http.Response) (*TitleSearchResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &TitleSearchResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseGetTitleResponse parses an HTTP response from a GetTitleWithResponse call
func ParseGetTitleResponse(rsp *http.Response) (*GetTitleResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetTitleResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}