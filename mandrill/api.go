package mandrill

import (
	"context"
	"encoding/json"
	"errors"
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
	InboundApi *inboundApi
}

func New(apiKey string) *API {
	api := &API{
		Key:    apiKey,
		client: http.DefaultClient,
	}

	api.PingApi = newPingApi(api)
	api.MessageApi = newMessageApi(api)
	api.InboundApi = newInboundApi(api)

	return api
}

type NormalResponse struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Name    string `json:"name"`
	Message string `json:"message"`
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

	if resp.StatusCode() < 200 || resp.StatusCode() > 300 {
		normalResp := NormalResponse{}
		if err := json.Unmarshal(resp.Body(), &normalResp); err != nil {
			return err
		}

		if normalResp.Status == "error" {
			return errors.New(normalResp.Message)
		}
	}

	fmt.Println(string(resp.Body()))

	return err
}
