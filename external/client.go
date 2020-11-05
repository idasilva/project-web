package external

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	wrap "github.com/pkg/errors"
	"net/http"
)

type Client struct {
	*http.Client
}

func (c *Client) Request(r interface{}) (*http.Response, error) {

	reqBytes := new(bytes.Buffer)
	err := json.NewEncoder(reqBytes).Encode(r)
	if err != nil {
		errInvalidEncode := errors.New("was not possible enconde response body")
		return nil, wrap.Wrap(err, errInvalidEncode.Error())

	}
	var UrlResult string
	var httStatus string

	switch r.(type) {
	case UrlImage:
		UrlResult = fmt.Sprintf("%s%s", "http://34.86.14.11:31582", endpoint3)
		httStatus = http.MethodPost

	case nil:

		UrlResult = fmt.Sprintf("%s%s", "http://34.86.14.11:31582", endpoint1)
		httStatus = http.MethodGet
	default:
		UrlResult = fmt.Sprintf("%s%s", "http://34.86.14.11:31582", endpoint2)
		httStatus = http.MethodPost

	}
	fmt.Println("ulr request:", UrlResult)

	request, err := http.NewRequest(
		httStatus,
		UrlResult,
		reqBytes,
	)
	request.Header.Set("Accept", "application/json; charset=utf-8")

	if err != nil {
		return nil, err
	}
	resp, err := c.Client.Do(request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

//NewClient return a new client instance
func NewClient() *Client {
	return &Client{
		&http.Client{},
	}
}
