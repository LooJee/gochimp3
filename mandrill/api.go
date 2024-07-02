package mandrill

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
)

const (
	ApiHost = "https://mandrillapp.com/api/1.0"
)

// API represents the origin of the API
type API struct {
	Key    string
	client *http.Client

	PingApi    *PingApi
	MessageApi *MessageApi
}

func New(apiKey string) *API {
	api := &API{
		Key:    apiKey,
		client: http.DefaultClient,
	}

	api.PingApi = newPingApi(api)
	api.MessageApi = newMessageApi(api)

	return api
}

func (api *API) Post(ctx context.Context, path string, body any, result any) error {
	request := resty.NewWithClient(api.client).NewRequest().SetContext(ctx)
	if body != nil {
		request.SetBody(body)
	}

	if result != nil {
		request.SetResult(result)
	}

	resp, err := request.Post(ApiHost + path)
	if err != nil {
		return err
	}

	fmt.Println(string(resp.Body()))

	return err
}
