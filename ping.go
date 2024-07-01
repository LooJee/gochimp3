package gochimp3

import "net/http"

const (
	ping_path = "/ping"
)

type PingResponse struct {
	HealthStatus string `json:"health_status"`
}

func (api *API) Ping() (*PingResponse, error) {
	response := &PingResponse{}

	return response, api.Request(http.MethodGet, ping_path, nil, nil, response)
}
