package mandrill

import (
	"context"
)

const (
	inboundDomainsPath = "/inbound/domains"
)

type inboundApi struct {
	api *API
}

func newInboundApi(api *API) *inboundApi {
	return &inboundApi{api: api}
}

type InboundDomainsReq struct {
	Key string `json:"key"`
}

type InboundDomainsRespItem struct {
	Domain    string `json:"domain"`
	CreatedAt string `json:"created_at"`
	ValidMx   bool   `json:"valid_mx"`
}

type InboundDomainsResp []InboundDomainsRespItem

func (api *inboundApi) Domains(ctx context.Context) (result InboundDomainsResp, err error) {
	req := InboundDomainsReq{Key: api.api.Key}

	err = api.api.Post(ctx, inboundDomainsPath, &req, &result)

	return
}
