package mandrill

import "context"

const (
	sendMsgPath = "/messages/send"
)

type SendMessageToItem struct {
	Email string `json:"email"`
	Name  string `json:"name,omitempty"`
	Type  string `json:"type"` //the header type to use for the recipient, defaults to "to" if not provided Possible values: "to", "cc", or "bcc".
}

type SendMessageItem struct {
	Text      string              `json:"text"`
	Subject   string              `json:"subject"`
	FromEmail string              `json:"from_email"`
	FromName  string              `json:"from_name,omitempty"`
	To        []SendMessageToItem `json:"to"`
}

type SendMessageReq struct {
	Key     string          `json:"key"`
	Message SendMessageItem `json:"message"`
	Async   bool            `json:"async,omitempty"`
	IpPool  string          `json:"ip_pool,omitempty"`
	SendAt  string          `json:"send_at,omitempty"`
}

type SendMessageRespItem struct {
	Email        string `json:"email"`
	Status       string `json:"status"`
	RejectReason string `json:"reject_reason"`
	QueuedReason string `json:"queued_reason"`
	Id           string `json:"_id"`
}

type SendMessageResp []SendMessageRespItem

type MessageApi struct {
	api *API
}

func newMessageApi(api *API) *MessageApi {
	return &MessageApi{api: api}
}

func (api *MessageApi) Send(ctx context.Context, req *SendMessageReq) (resp SendMessageResp, err error) {
	req.Key = api.api.Key

	err = api.api.Post(ctx, sendMsgPath, req, &resp)

	return resp, err
}
