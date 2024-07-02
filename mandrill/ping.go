package mandrill

import "context"

const (
	pingPath = "/users/ping"
)

type PingApi struct {
	api *API
}

func newPingApi(api *API) *PingApi {
	return &PingApi{api: api}
}

type PingReq struct {
	Key string `json:"key"`
}

func (api *PingApi) Ping(ctx context.Context) error {
	req := PingReq{Key: api.api.Key}

	return api.api.Post(ctx, pingPath, &req, nil)
}
